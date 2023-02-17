package routes

import (
	"github.com/gin-gonic/gin"
)

func addIndexRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "这是首页"})
	})
}
