package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"mio-init/controller"
	_ "mio-init/docs"
	"mio-init/logger"
	"mio-init/middleware"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.Cors, middleware.RefreshToken)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("/api")

	routerV1(apiGroup)

	r.NoRoute(func(c *gin.Context) {
		controller.ResponseError(c, controller.ErrorNotFound)
	})

	return r
}
