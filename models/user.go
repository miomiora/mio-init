package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserAccount  string  `json:"user_account" gorm:"size:256;comment:用户登录账号" binding:"required,min=4"`
	UserName     *string `json:"user_name,omitempty" gorm:"size:256;comment:用户名"`
	AvatarUrl    *string `json:"avatar_url,omitempty" gorm:"size:1024;comment:用户头像地址"`
	Gender       bool    `json:"gender" gorm:"comment:用户性别"`
	UserPassword string  `json:"user_password,omitempty" gorm:"size:256;comment:用户密码" binding:"required,min=8"`
	Phone        *string `json:"phone,omitempty" gorm:"size:256;comment:用户手机"`
	Email        *string `json:"email,omitempty" gorm:"size:256;comment:用户邮箱"`
	UserStatus   uint    `json:"user_status"  gorm:"size:2;comment:用户状态 0 表示正常"`
	Role         bool    `json:"role" gorm:"comment:0普通用户 1管理员"`
}

type UserDTO struct {
	ID          uint    `json:"id"`
	UserAccount string  `json:"user_account"`
	UserName    *string `json:"user_name"`
	Gender      bool    `json:"gender"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
}

type UserRegister struct {
	UserAccount   string `json:"user_account" binding:"required,min=4"`
	UserPassword  string `json:"user_password" binding:"required,min=8"`
	CheckPassword string `json:"check_password" binding:"required,eqfield=UserPassword"`
}
