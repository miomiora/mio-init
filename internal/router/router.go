package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	_ "mio-init/docs"
	"mio-init/internal/ctrls"
	middleware2 "mio-init/internal/middleware"
	"mio-init/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware2.Cors, middleware2.RefreshToken)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("/api")

	routerV1(apiGroup)

	r.NoRoute(func(c *gin.Context) {
		ctrls.ResponseError(c, ctrls.ErrorNotFound)
	})

	return r
}
