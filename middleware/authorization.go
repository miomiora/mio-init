package middleware

import (
	"github.com/gin-gonic/gin"
	"user-center/controllers"
	"user-center/utils"
)

func Authorization(c *gin.Context) {
	isValid, err := utils.ValidToken(c, controllers.Conn)
	if isValid {
		c.Next()
	}
	if err != nil {
		c.JSON(500, gin.H{
			"message": "读取Token用户失败！" + err.Error(),
			"data":    nil,
		})
		c.Abort()
	}
	c.Abort()
}
