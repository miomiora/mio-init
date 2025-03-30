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
