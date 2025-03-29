package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mio-init/dao/redis"
	"mio-init/internal/controller"
	"mio-init/util"
	"strconv"
)

var token = ""

func RefreshToken(c *gin.Context) {
	// 从请求头中获取Token, 没有token就直接返回
	token = c.GetHeader(util.TokenHeader)
	if token == util.TokenEmpty {
		c.Next()
	}
	redis.RefreshToken(token)
}

func AuthToken(c *gin.Context) {
	res, err := redis.CheckTokenExist(token)
	if err != nil {
		zap.L().Info("[middleware token AuthToken] token is not exist ", zap.Error(err))
		controller.ResponseError(c, controller.ErrorNotLogin)
		c.Abort()
		return
	}

	id, ok := res[0].(string)
	if !ok {
		zap.L().Info("[middleware token AuthToken] valid user_id failed ", zap.Error(err))
		controller.ResponseError(c, controller.ErrorNotLogin)
		c.Abort()
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		zap.L().Info("[middleware token AuthToken] valid user_role failed ", zap.Error(err))
		controller.ResponseError(c, controller.ErrorNotLogin)
		c.Abort()
		return
	}
	role, ok := res[1].(string)
	if !ok {
		zap.L().Info("[middleware token AuthToken] valid user_id failed ", zap.Error(err))
		controller.ResponseError(c, controller.ErrorNotLogin)
		c.Abort()
		return
	}
	userRole, err := strconv.Atoi(role)
	if err != nil {
		zap.L().Info("[middleware token AuthToken] valid user_role failed ", zap.Error(err))
		controller.ResponseError(c, controller.ErrorNotLogin)
		c.Abort()
		return
	}
	c.Set(util.KeyUserId, userId)
	c.Set(util.KeyUserRole, userRole)
	c.Next()
}

func AuthAdmin(c *gin.Context) {
	value, exist := c.Get(util.KeyUserRole)
	if !exist {
		zap.L().Warn("[middleware token AuthAdmin] get user_role error ")
		controller.ResponseError(c, controller.ErrorNotLogin)
		c.Abort()
		return
	}

	userRole, ok := value.(int)
	if !ok {
		zap.L().Warn("[middleware token AuthAdmin] get user error ")
		controller.ResponseError(c, controller.ErrorNotLogin)
		c.Abort()
		return
	}

	if userRole != util.RoleAdmin {
		zap.L().Warn("[middleware token AuthAdmin] auth user role is not admin ")
		controller.ResponseError(c, controller.ErrorNoAuth)
		c.Abort()
		return
	}

	c.Next()
}
