package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	controller2 "mio-init/internal/controller"
	"mio-init/model"
	"mio-init/service"
	"mio-init/util"
	"strconv"
)

type userController struct {
}

var User = new(userController)

// Login
// @Summary 用户登录
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param object body model.UserDTOLogin true "登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/login [post]
func (userController) Login(c *gin.Context) {
	// 1、校验参数
	u := new(model.UserDTOLogin)
	if err := c.ShouldBindJSON(u); err != nil {
		// 请求参数有误
		zap.L().Warn("[controller userController Login] login with invalid param ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	// 2、业务处理
	data, err := service.User.Login(u)
	if err != nil {
		zap.L().Error("[controller userController Login] login failed ", zap.String("Account", u.Account), zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			controller2.ResponseErrorWithMsg(c, controller2.ErrorInvalidParams, "用户名或密码错误！")
			return
		}
		controller2.ResponseError(c, controller2.ErrorServerBusy)
		return
	}
	// 3、返回响应
	controller2.ResponseOK(c, data)
}

// Logout
// @Summary 用户登出
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/logout [post]
func (userController) Logout(c *gin.Context) {
	token := c.GetHeader(util.TokenHeader)
	err := service.User.Logout(token)
	if err != nil {
		zap.L().Warn("[controller userController Logout] logout failed ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorServerBusy)
		return
	}
	controller2.ResponseOK(c, nil)
}

// Register
// @Summary 用户注册
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param object body model.UserDTORegister true "注册参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/register [post]
func (userController) Register(c *gin.Context) {
	// 1、参数校验
	u := new(model.UserDTORegister)
	if err := c.ShouldBindJSON(u); err != nil {
		// 请求参数有误
		zap.L().Warn("[controller userController Register] register with invalid param ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	// 2、业务处理
	if err := service.User.Register(u); err != nil {
		zap.L().Warn("[controller userController Register] register failed ", zap.Error(err))
		if errors.Is(err, service.ErrorUserExist) {
			controller2.ResponseErrorWithMsg(c, controller2.ErrorInvalidParams, err.Error())
			return
		}
		controller2.ResponseError(c, controller2.ErrorServerBusy)
		return
	}
	// 3、返回响应
	controller2.ResponseOK(c, nil)
}

// GetLoginUser
// @Summary 获取当前登录用户
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/get/login [get]
func (userController) GetLoginUser(c *gin.Context) {
	// 根据 user id 返回对应的用户
	userId, err := getUserId(c)
	if err != nil {
		zap.L().Warn("[controller userController GetLoginUser] get userId error ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorNotLogin)
		return
	}
	data, err := service.User.GetLoginUser(userId)
	if err != nil {
		zap.L().Warn("[controller userController GetLoginUser] get login user failed ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorForbidden)
		return
	}
	controller2.ResponseOK(c, data)
}

// UpdateBySelf
// @Summary 当前用户更新自己的信息
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param object body model.UserDTOUpdateBySelf true "修改后的数据"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/update/my [post]
func (userController) UpdateBySelf(c *gin.Context) {
	// 验证参数
	u := new(model.UserDTOUpdateBySelf)
	err := c.ShouldBindJSON(u)
	if err != nil || u.Password != u.RePassword {
		zap.L().Warn("[controller userController UpdateBySelf] update by self with invalid param ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	// 验证是否本人
	userId, err := getUserId(c)
	if err != nil || userId != u.UserId {
		zap.L().Warn("[controller userController UpdateBySelf] get userId error ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorNotLogin)
		return
	}
	// 业务
	if err = service.User.UpdateBySelf(u); err != nil {
		zap.L().Warn("[controller userController UpdateBySelf] update by self failed ", zap.Error(err))
		if errors.Is(err, service.ErrorUserExist) {
			controller2.ResponseErrorWithMsg(c, controller2.ErrorInvalidParams, err.Error())
			return
		}
	}
	controller2.ResponseOK(c, nil)
}

// GetUserVOByUserId
// @Summary 根据userId查找用户视图
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param userId query string true "需要查找的用户id"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/get/vo [get]
func (userController) GetUserVOByUserId(c *gin.Context) {
	value := c.Query(util.KeyUserId)
	if value == "" {
		zap.L().Warn("[controller userController GetUserVOByUserId] query userId failed ")
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	userId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		zap.L().Warn("[controller userController GetUserVOByUserId] parse userId failed ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	data, err := service.User.GetUserVOByUserId(userId)
	if err != nil {
		zap.L().Warn("[controller userController GetUserVOByUserId] get user vo by userId error ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorServerBusy)
		return
	}
	controller2.ResponseOK(c, data)

}

// GetUserVOList
// @Summary 获取用户视图列表
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param object body model.ListParams true "分页查询需要的参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/list/page/vo [post]
func (userController) GetUserVOList(c *gin.Context) {
	params := new(model.ListParams)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Warn("[controller userController GetUserVOList] get user vo list with invalid param ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	data, err := service.User.GetUserVOList(params)
	if err != nil {
		zap.L().Warn("[controller userController GetUserVOList] get user vo list failed ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorServerBusy)
		return
	}
	controller2.ResponseOK(c, data)
}

// AddUser
// @Summary 管理员添加用户
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param object body model.UserDTOAdd true "新用户的数据"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/add [post]
func (userController) AddUser(c *gin.Context) {
	// 1、参数校验
	u := new(model.UserDTOAdd)
	if err := c.ShouldBindJSON(u); err != nil {
		// 请求参数有误
		zap.L().Error("[controller userController AddUser] add user with invalid param ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	// 2、业务处理
	if err := service.User.AddUser(u); err != nil {
		zap.L().Error("[controller userController AddUser] add user failed ", zap.Error(err))
		if errors.Is(err, service.ErrorUserExist) {
			controller2.ResponseErrorWithMsg(c, controller2.ErrorInvalidParams, err.Error())
			return
		}
		controller2.ResponseError(c, controller2.ErrorServerBusy)
		return
	}
	// 3、返回响应
	controller2.ResponseOK(c, nil)
}

// DeleteUserByUserId
// @Summary 管理员根据userId删除用户
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param userId query string true "需要删除的userId"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/delete [post]
func (userController) DeleteUserByUserId(c *gin.Context) {
	value := c.Query(util.KeyUserId)
	if value == "" {
		zap.L().Warn("[controller userController DeleteUserByUserId] query userId failed ")
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	userId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		zap.L().Warn("[controller userController DeleteUserByUserId] parse userId failed ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	if err = service.User.DeleteUserByUserId(userId); err != nil {
		zap.L().Warn("[controller userController DeleteUserByUserId] delete user by userId failed ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorServerBusy)
		return
	}
	controller2.ResponseOK(c, nil)
}

// UpdateUserByAdmin
// @Summary 管理员根据userId更新用户信息
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param object body model.UserDTOUpdateByAdmin true "需要更新的用户信息"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/update [post]
func (userController) UpdateUserByAdmin(c *gin.Context) {
	u := new(model.UserDTOUpdateByAdmin)
	err := c.ShouldBindJSON(u)
	if err != nil || u.Password != u.RePassword {
		zap.L().Warn("[controller userController UpdateUserByAdmin] update user by admin with invalid param ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	// 业务
	err = service.User.UpdateUserByAdmin(u)
	if err != nil {
		zap.L().Warn("[controller userController UpdateUserByAdmin] update user by admin failed ", zap.Error(err))
		if errors.Is(err, service.ErrorUserExist) {
			controller2.ResponseErrorWithMsg(c, controller2.ErrorInvalidParams, err.Error())
			return
		}
	}
	controller2.ResponseOK(c, nil)
}

// GetUserByUserId
// @Summary 管理员根据userId查询用户完整信息
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param userId query string true "需要查询的userId"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/get [get]
func (userController) GetUserByUserId(c *gin.Context) {
	value := c.Query(util.KeyUserId)
	if value == "" {
		zap.L().Warn("[controller userController GetUserVOByUserId] query userId failed ")
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	userId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		zap.L().Warn("[controller userController GetUserByUserId] parse userId failed ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	data, err := service.User.GetUserByUserId(userId)
	if err != nil {
		zap.L().Warn("[controller userController GetUserByUserId] get user by userId failed ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorServerBusy)
		return
	}
	controller2.ResponseOK(c, data)
}

// GetUserList
// @Summary 管理员根据查询用户完整信息列表
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param object body model.ListParams true "分页查询所需要的参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/list/page [post]
func (userController) GetUserList(c *gin.Context) {
	params := new(model.ListParams)
	err := c.ShouldBindJSON(params)
	if err != nil {
		zap.L().Warn("[controller userController GetUserList] get user list with invalid param ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorInvalidParams)
		return
	}
	data, err := service.User.GetUserList(params)
	if err != nil {
		zap.L().Warn("[controller userController GetUserList] get user list failed ", zap.Error(err))
		controller2.ResponseError(c, controller2.ErrorServerBusy)
		return
	}
	controller2.ResponseOK(c, data)
}
