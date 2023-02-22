package routes

import (
	"github.com/gin-gonic/gin"
	"user-center/api"
	"user-center/middlewares"
)

func addUserRouter(apiGroup *gin.RouterGroup) {
	userRouter := apiGroup.Group("user")
	{
		/*
			login		登录
			register	注册
			logout		登出
		*/
		userRouter.POST("login", api.UserLogin)
		userRouter.POST("register", api.UserRegister)
		userRouter.POST("logout", api.UserLogout)

		/*
			普通用户都可访问
			current		根据Token获取的id进行查找用户
			update		普通用户修改自己的信息
			change		普通用户修改自己的密码
		*/
		userRouter.GET("current", middlewares.AuthUser, api.GetCurrentUser)
		userRouter.PUT("update", middlewares.AuthUser, api.UpdateUserBySelf)
		userRouter.PUT("change", middlewares.AuthUser, api.ChangePasswordBySelf)

		/*
			下列需要管理员权限
			list/:num/:page		获取全部的用户(分页)
			search/:id			获取单个用户
			delete/:id			删除指定用户
			update/:id			更新指定用户
			change/:id			修改指定用户的密码
		*/
		userRouter.GET("list/:num/:page", middlewares.AuthAdmin, api.GetUserList)
		userRouter.GET("search/:id", middlewares.AuthAdmin, api.GetUserById)
		userRouter.POST("delete/:id", middlewares.AuthAdmin, api.DeleteUserById)
		userRouter.PUT("update/:id", middlewares.AuthAdmin, api.UpdateUserById)
		userRouter.PUT("change/:id", middlewares.AuthAdmin, api.ChangePasswordById)
	}
}
