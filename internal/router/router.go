package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	_ "mio-init/docs"
	"mio-init/internal/middleware"
	"mio-init/logger"
	"mio-init/util"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.Cors)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("/api")

	routerV1(apiGroup)

	r.NoRoute(func(c *gin.Context) {
		util.ResponseError(c, util.ErrorNotFound)
	})

	return r
}
