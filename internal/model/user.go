package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserId   int64  `json:"userId" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Account  string `json:"account" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

type UserCreateReq struct {
	Account  string `json:"account" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}
