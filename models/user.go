package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserAccount  string  `json:"user_account" gorm:"size:256;comment:用户登录账号;unique" binding:"required,min=4"`
	UserName     *string `json:"user_name,omitempty" gorm:"size:256;comment:用户名"`
	AvatarUrl    *string `json:"avatar_url,omitempty" gorm:"size:1024;comment:用户头像地址"`
	Gender       uint    `json:"gender" gorm:"comment:用户性别;size:2"`
	UserPassword string  `json:"user_password,omitempty" gorm:"size:256;comment:用户密码" binding:"required,min=8"`
	Phone        *string `json:"phone,omitempty" gorm:"size:256;comment:用户手机"`
	Email        *string `json:"email,omitempty" gorm:"size:256;comment:用户邮箱"`
	UserStatus   uint    `json:"user_status"  gorm:"size:2;comment:用户状态 0 表示正常"`
	Role         uint    `json:"role" gorm:"comment:0普通用户 1管理员;size:2"`
}

type UserDTO struct {
	ID          uint    `json:"id" binding:"required"`
	UserAccount string  `json:"user_account" binding:"required,min=4"`
	UserName    *string `json:"user_name"`
	Gender      uint    `json:"gender"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	AvatarUrl   *string `json:"avatar_url"`
	Role        uint    `json:"role"`
	UserStatus  uint    `json:"user_status"`
}

type UserRegister struct {
	UserAccount   string `json:"user_account" binding:"required,min=4"`
	UserPassword  string `json:"user_password" binding:"required,min=8"`
	CheckPassword string `json:"check_password" binding:"required,eqfield=UserPassword"`
}

type UserChangePassword struct {
	ID            uint   `json:"id" binding:"required"`
	UserPassword  string `json:"user_password" binding:"required,min=8"`
	CheckPassword string `json:"check_password" binding:"required,eqfield=UserPassword"`
}
