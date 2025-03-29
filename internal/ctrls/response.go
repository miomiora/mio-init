package ctrls

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorCode struct {
	HttpStatus int
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

// 定义错误码
var (
	ErrorInvalidParams = ErrorCode{http.StatusForbidden, 40000, "请求参数错误"}
	ErrorNotLogin      = ErrorCode{http.StatusUnauthorized, 40100, "未登录"}
	ErrorNoAuth        = ErrorCode{http.StatusUnauthorized, 400101, "无权限"}
	ErrorForbidden     = ErrorCode{http.StatusForbidden, 400300, "拒绝访问"}
	ErrorNotFound      = ErrorCode{http.StatusNotFound, 400400, "请求数据不存在"}
	ErrorServerBusy    = ErrorCode{http.StatusInternalServerError, 500000, "系统内部繁忙"}
	ErrorRedisBusy     = ErrorCode{http.StatusInternalServerError, 500001, "Redis繁忙"}
	ErrorMysqlBusy     = ErrorCode{http.StatusInternalServerError, 500002, "Mysql繁忙"}
)

func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Message: "ok",
		Code:    0,
		Data:    data,
	})
}

func ResponseOKWithMsg(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, &Response{
		Message: message,
		Code:    0,
		Data:    data,
	})
}

func ResponseError(c *gin.Context, errorCode ErrorCode) {
	c.JSON(errorCode.HttpStatus, &Response{
		Message: errorCode.Message,
		Code:    errorCode.Code,
		Data:    nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, errorCode ErrorCode, message string) {
	c.JSON(errorCode.HttpStatus, &Response{
		Message: message,
		Code:    errorCode.Code,
		Data:    nil,
	})
}

func ResponseErrorWithHttpCode(c *gin.Context, errorCode ErrorCode, httpCode int) {
	c.JSON(httpCode, &Response{
		Message: errorCode.Message,
		Code:    errorCode.Code,
		Data:    nil,
	})
}
