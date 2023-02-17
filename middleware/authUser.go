package middleware

import (
	"fmt"
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
	role, err := utils.ValidToken(c, api.Conn, api.DB)
	if err == nil {
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
			c.Next()
		// 管理员
		case utils.ROLE_ADMIN:
			c.Next()
		default:
			c.Abort()
		}
	} else {
		fmt.Println("AuthUser err!!!" + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请登录",
			"data":    nil,
		})
		c.Abort()
	}
}
