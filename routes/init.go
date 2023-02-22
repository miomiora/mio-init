package routes

import (
	"github.com/gin-gonic/gin"
	"user-center/middlewares"
)

var R *gin.Engine

// 初始化gin
func init() {
	r := gin.Default()
	r.Use(middlewares.Cors)
	r.Use(middlewares.RefreshRedis)
	apiGroup := r.Group("api")
	addUserRouter(apiGroup)
	R = r
}
