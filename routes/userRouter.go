package routes

import (
	"github.com/gin-gonic/gin"
	"user-center/controllers"
)

func addUserRouter(r *gin.Engine) {
	userRouter := r.Group("user")
	{
		userRouter.POST("login", controllers.UserLogin)
		userRouter.POST("register", controllers.UserRegister)
	}
}
