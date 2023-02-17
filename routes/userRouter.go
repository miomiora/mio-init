package routes

import (
	"github.com/gin-gonic/gin"
	"user-center/api"
	"user-center/middleware"
)

func addUserRouter(r *gin.Engine) {
	userRouter := r.Group("user")
	{
		userRouter.POST("login", api.UserLogin)
		userRouter.POST("register", api.UserRegister)

		// 普通用户都可访问
		// 在请求login页面的时候，进行判断是否已经登录过，如果登录直接告知前端已经登录
		userRouter.GET("login", middleware.AuthUser, api.GetUserLoginPage)

		// 下列需要管理员权限
		userRouter.GET(":id", middleware.AuthAdmin, api.GetUserById)
		userRouter.GET("search", middleware.AuthAdmin, api.GetUserList)
		userRouter.POST("delete/:id", middleware.AuthAdmin, api.DeleteUserById)
	}
}
