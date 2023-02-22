package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-center/api"
	"user-center/utils"
)

//
// RefreshRedis
//  @Description: 全局中间件，判断有token就刷新，无token直接放行
//  @param c
//
func RefreshRedis(c *gin.Context) {
	// 从请求头中获取Token, 没有token就直接返回
	token := c.GetHeader("Authorization")
	if token == "" {
		c.Next()
	}
	tokenKey := utils.TokenPrefix + token
	// 在redis中查找是否Token是否存在
	do, err := api.Conn.Do("HGET", tokenKey, "id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
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
	_, err = api.Conn.Do("EXPIRE", tokenKey, utils.TokenTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "设置Token有效期失败！" + err.Error(),
			"data":    nil,
		})
		return
	}
}
