package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func ValidToken(c *gin.Context, conn redis.Conn) (bool, error) {
	// 从请求头中获取Token, 没有token就直接返回
	token := c.GetHeader("Authorization")
	if token == "" {
		return false, nil
	}
	tokenKey := "login:token:" + token
	// 从redis中进行查询
	ip, err := redis.String(conn.Do("HGET", tokenKey, "client_ip"))
	// 没查到就会报err
	if err != nil {
		return false, err
	}
	// 判断ip是否和当前客户端请求的ip一致
	if ip != c.ClientIP() {
		return false, nil
	}
	return true, nil
}
