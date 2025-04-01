package ctrls

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mio-init/internal/model"
	"mio-init/internal/service"
	"mio-init/util"
)

type userCtrl struct {
}

var User = new(userCtrl)

// Login
// @Summary 用户登录
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param object body model.UserLoginReq true "登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/login [post]
func (userCtrl) Login(c *gin.Context) {
	// 1、校验参数
	u := new(model.UserLoginReq)
	if err := c.ShouldBindJSON(u); err != nil {
		// 请求参数有误
		zap.L().Warn("[ctrls userCtrl Login] login with invalid param ", zap.Error(err))
		util.ResponseError(c, util.ErrorInvalidParams)
		return
	}
	// 2、业务处理
	data, err := service.User.Login(c.Request.Context(), u)
	if err != nil {
		zap.L().Error("[ctrls userCtrl Login] login failed ", zap.String("Account", u.Account), zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseErrorWithMsg(c, util.ErrorInvalidParams, "用户名或密码错误！")
			return
		}
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}

	c.SetCookie(util.KeyRefresh, data.RefreshToken, int(util.RefreshTokenExpire.Seconds()), "/", "", true, true)
	// 3、返回响应
	util.ResponseOK(c, data)
}

// Create
// @Summary 用户注册
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param object body model.UserCreateReq true "注册参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/register [post]
func (userCtrl) Create(c *gin.Context) {
	// 1、参数校验
	u := new(model.UserCreateReq)
	if err := c.ShouldBindJSON(u); err != nil {
		// 请求参数有误
		zap.L().Warn("[ctrls userCtrl Create] register with invalid param ", zap.Error(err))
		util.ResponseError(c, util.ErrorInvalidParams)
		return
	}
	// 2、业务处理
	if err := service.User.Create(c.Request.Context(), u); err != nil {
		zap.L().Warn("[ctrls userCtrl Create] register failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}
	// 3、返回响应
	util.ResponseOK(c, nil)
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
func (userCtrl) Logout(c *gin.Context) {
	accessToken := c.GetHeader(util.TokenHeader)
	refreshToken, err := c.Cookie(util.KeyRefresh)
	if err != nil {
		zap.L().Warn("[ctrls userCtrl Logout] Logout failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}

	err = service.User.Logout(c.Request.Context(), refreshToken, accessToken)
	if err != nil {
		zap.L().Warn("[ctrls userCtrl Logout] Logout failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}
	util.ResponseOK(c, nil)
}

// GetBySelf
// @Summary 用户信息
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/get/my [get]
func (userCtrl) GetBySelf(c *gin.Context) {
	user, err := service.User.GetByUserId(c.Request.Context(), util.GetUserIdByContext(c))
	if err != nil {
		zap.L().Warn("[ctrls userCtrl Logout] GetBySelf failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}
	util.ResponseOK(c, user)
}

// Update
// @Summary 修改用户信息
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param object body model.UserUpdateReq true "更新参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/update [post]
func (userCtrl) Update(c *gin.Context) {
	// 1、参数校验
	u := new(model.UserUpdateReq)
	if err := c.ShouldBindJSON(u); err != nil {
		// 请求参数有误
		zap.L().Warn("[ctrls userCtrl Create] Update with invalid param ", zap.Error(err))
		util.ResponseError(c, util.ErrorInvalidParams)
		return
	}
	// 2、业务处理
	if err := service.User.Update(c.Request.Context(), u, util.GetUserIdByContext(c)); err != nil {
		zap.L().Warn("[ctrls userCtrl Create] register failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}
	// 3、返回响应
	util.ResponseOK(c, nil)
}

// UpdatePassword
// @Summary 修改密码
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param object body model.UserUpdatePwdReq true "更新参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user/update/pwd [post]
func (userCtrl) UpdatePassword(c *gin.Context) {
	accessToken := c.GetHeader(util.TokenHeader)
	refreshToken, err := c.Cookie(util.KeyRefresh)
	userId := util.GetUserIdByContext(c)
	if err != nil {
		// 请求参数有误
		zap.L().Warn("[ctrls userCtrl Create] UpdatePassword with invalid param ", zap.Error(err))
		util.ResponseError(c, util.ErrorInvalidParams)
		return
	}
	// 1、参数校验
	u := new(model.UserUpdatePwdReq)
	if err := c.ShouldBindJSON(u); err != nil || userId != u.UserId {
		// 请求参数有误
		zap.L().Warn("[ctrls userCtrl Create] UpdatePassword with invalid param ", zap.Error(err))
		util.ResponseError(c, util.ErrorInvalidParams)
		return
	}
	// 2、业务处理
	if err := service.User.UpdatePwd(c.Request.Context(), u, accessToken, refreshToken); err != nil {
		zap.L().Warn("[ctrls userCtrl Create] UpdatePassword failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}
	// 3、返回响应
	util.ResponseOK(c, nil)
}

// GetByUserId
// @Summary 获取用户信息
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /user [get]
func (userCtrl) GetByUserId(c *gin.Context) {
	userId, err := util.GetUserIdByParam(c)
	if err != nil {
		zap.L().Warn("[ctrls userCtrl GetByUserId] GetByUserId with invalid param ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}
	user, err := service.User.GetByUserId(c.Request.Context(), userId)
	if err != nil {
		zap.L().Warn("[ctrls userCtrl GetByUserId] GetByUserId failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}

	util.ResponseOK(c, user)
}

// Delete
// @Summary 删除用户
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /delete [get]
func (userCtrl) Delete(c *gin.Context) {
	userId, err := util.GetUserIdByParam(c)
	if err != nil {
		zap.L().Warn("[ctrls userCtrl GetByUserId] Delete with invalid param ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}
	err = service.User.Delete(c.Request.Context(), userId)
	if err != nil {
		zap.L().Warn("[ctrls userCtrl GetByUserId] Delete failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}

	util.ResponseOK(c, nil)
}

// List
// @Summary 用户列表
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /list [get]
func (userCtrl) List(c *gin.Context) {
	page, err := util.GetIntByQuery(c, util.KeyPage)
	pageSize, err := util.GetIntByQuery(c, util.KeyPageSize)
	orderBy := c.Query(util.KeyOrderBy)
	if err != nil {
		zap.L().Warn("[ctrls userCtrl List] List with invalid param ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}
	users, count, err := service.User.List(c.Request.Context(), page, pageSize, orderBy)
	if err != nil {
		zap.L().Warn("[ctrls userCtrl List] List failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}

	util.ResponseOK(c, gin.H{
		"total":  count,
		"record": users,
	})
}
