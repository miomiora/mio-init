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
		zap.L().Warn("[ctrls userCtrl Logout] logout failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}

	err = service.User.Logout(c.Request.Context(), refreshToken, accessToken)
	if err != nil {
		zap.L().Warn("[ctrls userCtrl Logout] logout failed ", zap.Error(err))
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
		zap.L().Warn("[ctrls userCtrl Logout] logout failed ", zap.Error(err))
		util.ResponseError(c, util.ErrorServerBusy)
		return
	}
	util.ResponseOK(c, user)
}
