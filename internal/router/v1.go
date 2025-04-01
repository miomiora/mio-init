package router

import (
	"github.com/gin-gonic/gin"
	"mio-init/internal/ctrls"
	"mio-init/internal/middleware"
)

func routerV1(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	userRouter(v1)
	//postRouter(v1)
}

func userRouter(r *gin.RouterGroup) {
	user := r.Group("/user")

	// 登录注册
	user.POST("/login", ctrls.User.Login)
	user.POST("/register", ctrls.User.Create)

	// 以下操作需要登录
	user.Use(middleware.AuthToken)
	user.POST("/logout", ctrls.User.Logout)
	user.POST("/update", ctrls.User.Update)
	user.POST("/update/pwd", ctrls.User.UpdatePassword)
	user.GET("/get/:user_id", ctrls.User.GetByUserId)
	user.GET("/get/my", ctrls.User.GetBySelf)
	user.GET("/delete/:user_id", ctrls.User.Delete)
	user.GET("/list", ctrls.User.List)
}
