package middleware

import (
	"github.com/gin-gonic/gin"
	"mio-init/internal/repository"
	"mio-init/util"
	"strconv"
)

func AuthToken(c *gin.Context) {
	accessToken := c.GetHeader(util.TokenHeader)

	if accessToken == "" {
		util.ResponseError(c, util.ErrorNotLogin)
		c.Abort()
		return
	}

	// 2. 解析并验证 JWT
	claims, err := util.ParseToken(accessToken)
	if err != nil {
		newAccess, ok := tryRefreshToken(c)
		if !ok {
			util.ResponseError(c, util.ErrorNotLogin)
			c.Abort()
			return
		}

		// 3. 返回新 Token 给前端（提示重试）
		util.ResponseOKRetry(c, newAccess)
		c.Abort()
		return
	}

	// 3. （可选）检查黑名单（如 /logout 或敏感操作）
	if isSensitivePath(c.Request.URL.Path) {
		if exists, _ := repository.Cache.Exists(c.Request.Context(), util.GenBlackListKey(accessToken)); exists == 1 {
			util.ResponseError(c, util.ErrorNotLogin)
			c.Abort()
			return
		}
	}

	c.Set(util.KeyUserId, claims[util.KeyUserId].(string))
	c.Next()
}

func tryRefreshToken(c *gin.Context) (newAccess string, ok bool) {
	refreshToken, _ := c.Cookie(util.KeyRefresh)
	if refreshToken == "" {
		return "", false
	}

	// 验证并生成新 Token
	result, err := repository.Cache.Get(c, util.GenRefreshKey(refreshToken))
	if err != nil {
		return "", false
	}
	userId, _ := strconv.ParseInt(result, 10, 64)
	newAccess, newRefresh, _ := util.GenTokens(userId)

	err = repository.Cache.RefreshToken(c.Request.Context(), refreshToken, newRefresh, userId)
	if err != nil {
		return "", false
	}

	// 设置新 Cookie
	c.SetCookie(util.KeyRefresh, newRefresh, int(util.RefreshTokenExpire.Seconds()), "/", "", true, true)
	return newAccess, true
}

func isSensitivePath(path string) bool {
	sensitivePaths := map[string]bool{
		"/user/update/pwd": true,
		"/user/logout":     true,
	}
	return sensitivePaths[path]
}
