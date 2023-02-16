package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserAccount  string  `json:"user_account" gorm:"size:256"`
	UserName     string  `json:"user_name,omitempty" gorm:"size:256"`
	AvatarUrl    string  `json:"avatar_url,omitempty" gorm:"size:1024"`
	Gender       bool    `json:"gender"`
	UserPassword string  `json:"user_password,omitempty" gorm:"size:256"`
	Phone        *string `json:"phone,omitempty" gorm:"size:256"`
	Email        *string `json:"email,omitempty" gorm:"size:256"`
	UserStatus   uint    `json:"user_status"  gorm:"size:2"`
}

type UserDTO struct {
	UserName string  `json:"user_name"`
	Gender   bool    `json:"gender"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
}

type UserRegister struct {
	UserAccount   string `json:"user_account" binding:"required,min=4"`
	UserPassword  string `json:"user_password" binding:"required,min=8"`
	CheckPassword string `json:"check_password" binding:"required,eqfield=UserPassword"`
}
