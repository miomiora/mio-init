package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-center/api"
	"user-center/utils"
)

//
// AuthUser
//  @Description: token无效的情况下不放行, 在role为ROLE_USER或ROLE_ADMIN的时候才放行
//  @param c
//
func AuthUser(c *gin.Context) {
	utils.ValidToken(c, api.Conn, api.DB)
	role, exists := c.Get("userRole")
	// 不存在则意味着未存入任何内容到userRole中，服务器错误
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取用户权限错误！",
		})
		c.Abort()
	}
	switch role {
	// token验证失败
	case utils.ROLE_UNDEFINED:
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请登录",
		})
		c.Abort()
	// 普通用户
	case utils.ROLE_USER:
		c.Next()
	// 管理员
	case utils.ROLE_ADMIN:
		c.Next()
	default:
		c.Abort()
	}
}
