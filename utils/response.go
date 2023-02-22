package utils

type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Message     string      `json:"message"`
	Code        int         `json:"code"`
	Data        interface{} `json:"data"`
	Description *string     `json:"description"`
}

var ParamsError = ErrorCode{40000, "请求参数错误！"}
var NotLogin = ErrorCode{40100, "未登录!"}
var NoAuth = ErrorCode{400101, "无权限!"}
var ServerError = ErrorCode{500000, "服务器错误!"}
var RedisError = ErrorCode{500001, "Redis错误!"}
var MysqlError = ErrorCode{500002, "Mysql错误!"}

func ResponseOK(data interface{}, description ...string) *Response {
	var des *string
	if len(description) > 0 {
		des = &description[0]
	} else {
		des = nil
	}
	return &Response{
		Message:     "ok",
		Code:        0,
		Data:        data,
		Description: des,
	}
}

func ResponseError(errorCode ErrorCode, description ...string) *Response {
	var des *string
	if len(description) > 0 {
		des = &description[0]
	} else {
		des = nil
	}
	return &Response{
		Message:     errorCode.Message,
		Code:        errorCode.Code,
		Data:        nil,
		Description: des,
	}
}
