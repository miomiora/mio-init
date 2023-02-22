package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-center/api"
	"user-center/utils"
)

//
// AuthAdmin
//  @Description: token无效的情况下不放行, 仅在role为ROLE_ADMIN的时候放行
//  @param c
//
func AuthAdmin(c *gin.Context) {
	utils.ValidToken(c, api.Conn, api.DB)
	role, exists := c.Get("userRole")
	// 不存在则意味着未存入任何内容到userRole中，服务器错误
	if !exists {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError))
		c.Abort()
	}
	switch role {
	// token验证失败
	case utils.RoleUndefined:
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.NotLogin))
		c.Abort()
	// 普通用户
	case utils.RoleUser:
		c.JSON(http.StatusUnauthorized, utils.ResponseError(utils.NoAuth))
		c.Abort()
	// 管理员
	case utils.RoleAdmin:
		c.Next()
	default:
		c.Abort()
	}
}
