package middleware

import (
	"github.com/gin-gonic/gin"
	"user-center/controllers"
)

func RefreshRedis(c *gin.Context) {
	// 从请求头中获取Token, 没有token就直接返回
	token := c.GetHeader("Authorization")
	if token == "" {
		c.Next()
	}
	tokenKey := "login:token:" + token
	// 在redis中查找是否Token是否存在
	do, err := controllers.Conn.Do("HGET", tokenKey, "id")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "查找Token失败！" + err.Error(),
			"data":    nil,
		})
		return
	}
	// 没有查询到任何用户，直接放行
	if do == nil {
		c.Next()
	}
	// 查找到了用户，刷新token时间
	_, err = controllers.Conn.Do("EXPIRE", tokenKey, 600)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "设置Token有效期失败！" + err.Error(),
			"data":    nil,
		})
		return
	}
}
