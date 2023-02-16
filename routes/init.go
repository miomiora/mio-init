package routes

import (
	"github.com/gin-gonic/gin"
	"user-center/middleware"
)

var R *gin.Engine

// 初始化gin
func init() {
	r := gin.Default()
	r.Use(middleware.RefreshRedis)
	addIndexRouter(r)
	addUserRouter(r)

	R = r
}
