package routes

import (
	"github.com/gin-gonic/gin"
	"user-center/middlewares"
)

var R *gin.Engine

// 初始化gin
func init() {
	r := gin.Default()
	// 全局中间件
	r.Use(middlewares.Cors, middlewares.RefreshRedis)
	apiGroup := r.Group("api")
	addUserRouter(apiGroup)
	R = r
}
