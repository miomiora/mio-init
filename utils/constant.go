package utils

const (
	/*
		Token验证结果

		-1:验证失败
		 0:验证为用户
		 1:验证为管理员
	*/
	ROLE_UNDEFINED = -1
	ROLE_USER      = 0
	ROLE_ADMIN     = 1

	/*
		Token在redis中存储的Key前缀
	*/
	TOKEN_PREIX = "login:token:"

	/*
		Token的存在时间 秒
	*/
	TOKEN_TIMEOUT = 6000
)
