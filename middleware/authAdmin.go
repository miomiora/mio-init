package middleware

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
	role, err := utils.ValidToken(c, api.Conn, api.DB)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请登录" + err.Error(),
			"data":    nil,
		})
		c.Abort()
	}
	switch role {
	// token验证失败
	case utils.ROLE_UNDEFINED:
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请登录",
			"data":    nil,
		})
		c.Abort()
	// 普通用户
	case utils.ROLE_USER:
		c.JSON(http.StatusForbidden, gin.H{
			"message": "权限不足！需要管理员权限！",
		})
		c.Abort()
	// 管理员
	case utils.ROLE_ADMIN:
		c.Next()
	default:
		c.Abort()
	}
}
