package router

import (
	"github.com/gin-gonic/gin"
	"mio-init/controller"
	"mio-init/middleware"
)

func routerV1(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	userRouter(v1)
	postRouter(v1)
}

func userRouter(r *gin.RouterGroup) {
	user := r.Group("/user")

	// 登录注册
	user.POST("/login", controller.User.Login)
	user.POST("/register", controller.User.Register)

	// 以下操作需要登录
	user.Use(middleware.AuthToken)
	user.POST("/logout", controller.User.Logout)
	user.GET("/get/login", controller.User.GetLoginUser)
	user.POST("/update/my", controller.User.UpdateBySelf)
	user.GET("/get/vo", controller.User.GetUserVOByUserId)
	user.POST("/list/page/vo", controller.User.GetUserVOList)

	// 仅管理员
	user.Use(middleware.AuthAdmin)
	user.POST("/add", controller.User.AddUser)
	user.POST("/delete", controller.User.DeleteUserByUserId)
	user.POST("/update", controller.User.UpdateUserByAdmin)
	user.GET("/get", controller.User.GetUserByUserId)
	user.POST("/list/page", controller.User.GetUserList)
}

func postRouter(r *gin.RouterGroup) {
	post := r.Group("/post")

	// 以下操作需要登录
	post.Use(middleware.AuthToken)
	post.POST("/new", controller.Post.InsertPost)
	post.POST("/update/my", controller.Post.UpdateBySelf)
	post.GET("/get/vo", controller.Post.GetPostVOByPostId)
	post.POST("/my", controller.Post.GetMyPostVOList)
	post.POST("/list/page/vo", controller.Post.GetPostVOList)

	// 仅管理员
	post.Use(middleware.AuthAdmin)
	post.POST("/add", controller.Post.AddPost)
	post.POST("/delete", controller.Post.DeletePostByPostId)
	post.POST("/update", controller.Post.UpdatePostByAdmin)
	post.GET("/get", controller.Post.GetPostByPostId)
	post.POST("/list/page", controller.Post.GetPostList)
}
