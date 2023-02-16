package routes

import (
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

// 初始化gin
func init() {
	r := gin.Default()
	addIndexRouter(r)
	addUserRouter(r)
	R = r
}
