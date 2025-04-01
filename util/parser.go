package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserIdByContext(c *gin.Context) int64 {
	inter, _ := c.Get(KeyUserId)
	str := inter.(string)
	userId, _ := strconv.ParseInt(str, 10, 64)
	return userId
}

func GetUserIdByQuery(c *gin.Context) (int64, error) {
	value := c.Query(KeyUserId)
	return strconv.ParseInt(value, 10, 64)
}

func GetUserIdByParam(c *gin.Context) (int64, error) {
	value := c.Param(KeyUserId)
	return strconv.ParseInt(value, 10, 64)
}

func GetIntByQuery(c *gin.Context, key string) (int, error) {
	value := c.Query(key)
	return strconv.Atoi(value)
}
