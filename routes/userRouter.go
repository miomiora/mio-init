package routes

import (
	"github.com/gin-gonic/gin"
	"user-center/controllers"
	"user-center/middleware"
)

func addUserRouter(r *gin.Engine) {
	userRouter := r.Group("user")
	{
		// 防止有人使用postman直接跳过前端的get请求，直接给后端发送post登录请求，造成多次插入不同的token到redis中 虽然人家不带上请求头一样可以用postman^^
		userRouter.POST("login", controllers.UserLogin)
		// 在请求login页面的时候，进行判断是否已经登录过，如果登录直接告知前端
		userRouter.GET("login", controllers.GetLoginPage)
		userRouter.POST("register", controllers.UserRegister)
		userRouter.GET(":id", middleware.Authorization, controllers.GetUserById)
	}
}
