package utils

const (
	/*
		Token验证结果

		-1:验证失败
		 0:验证为用户
		 1:验证为管理员
	*/
	RoleUndefined = -1
	RoleUser      = 0
	RoleAdmin     = 1

	/*
		Token在redis中存储的Key前缀
	*/
	TokenPrefix = "login:token:"

	/*
		Token的存在时间 秒
	*/
	TokenTimeout = 6000
)
