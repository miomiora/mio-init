package model

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model

	UserId      int64   `json:"user_id" gorm:"not null"`
	Account     string  `json:"account" gorm:"not null;unique"`
	Password    string  `json:"password" gorm:"not null"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	Description *string `json:"description"`
	Gender      bool    `json:"gender"`
	UserRole    uint8   `json:"user_role" gorm:"comment:0-普通用户 1-管理员;size:2"`
}

// UserDTOLogin 用户登录所需要绑定的参数
type UserDTOLogin struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserDTORegister 用户注册所需要绑定的参数
type UserDTORegister struct {
	Account    string `json:"account" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// UserDTOAdd 用户登录所需要绑定的参数
type UserDTOAdd struct {
	Account    string `json:"account" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// UserDTOUpdateBySelf 用户更新参数
type UserDTOUpdateBySelf struct {
	UserId      int64   `json:"user_id,string" binding:"required"`
	Account     string  `json:"account" binding:"required"`
	Password    string  `json:"password"`
	RePassword  string  `json:"re_password"`
	Description *string `json:"description"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	Gender      bool    `json:"gender"`
}

// UserDTOUpdateByAdmin 用户更新参数
type UserDTOUpdateByAdmin struct {
	UserId      int64   `json:"user_id,string" binding:"required"`
	Account     string  `json:"account" binding:"required"`
	Password    string  `json:"password"`
	RePassword  string  `json:"re_password"`
	Description *string `json:"description"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	Gender      bool    `json:"gender"`
	UserRole    uint8   `json:"user_role"`
}

// UserVO 登录成功返回给前端展示的用户数据
type UserVO struct {
	UserId      int64   `json:"user_id,string"`
	Account     string  `json:"account"`
	Token       *string `json:"token,omitempty"`
	Description *string `json:"description,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
	Gender      bool    `json:"gender"`
	UserRole    uint8   `json:"user_role"`
}
