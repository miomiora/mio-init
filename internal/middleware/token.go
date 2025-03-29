package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"mio-init/util"
)

func RefreshToken(c *gin.Context) {
	// 1. 获取 Refresh Token
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// 2. 验证 Refresh Token
	claims, err := util.ParseToken(req.RefreshToken)
	if err != nil || claims["type"] != "refresh" {
		c.JSON(401, gin.H{"error": "Invalid refresh token"})
		return
	}

	// 3. 检查 Redis 中是否存在
	userID, err := redis.Client.Get(c, "refresh:"+req.RefreshToken).Uint64()
	if err != nil {
		c.JSON(401, gin.H{"error": "Refresh token expired"})
		return
	}

	// 4. 生成新的 Token 对
	newAccess, newRefresh, err := util.GenerateTokens(uint(userID))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// 5. 更新 Redis 中的 Refresh Token
	_ = redis.Client.Del(c.Request.Context(), "refresh:"+req.RefreshToken)
	_ = redis.Client.Set(c, "refresh:"+newRefresh, userID, jwt.RefreshTokenExpire)

	c.JSON(200, gin.H{
		"access_token":  newAccess,
		"refresh_token": newRefresh,
	})
}
