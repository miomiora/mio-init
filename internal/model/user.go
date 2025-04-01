package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   int64  `json:"userId" gorm:"not null;index"`
	Name     string `json:"name" gorm:"not null;size:100"`
	Account  string `json:"account" gorm:"not null;uniqueIndex;size:50"`
	Password string `json:"password" gorm:"not null;size:128"`
}

type UserCreateReq struct {
	Account    string `json:"account" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

type UserLoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateReq struct {
	UserId int64  `json:"userId" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

type UserUpdatePwdReq struct {
	Account     string `json:"account" binding:"required"`
	Password    string `json:"password" binding:"required"`
	UserId      int64  `json:"userId"`
	NewPassword string `json:"newPassword" binding:"required"`
	RePassword  string `json:"rePassword" binding:"required,eqfield=NewPassword"`
}

type UserLoginRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserId       int64  `json:"userId"`
	Name         string `json:"name"`
	Account      string `json:"account"`
}

type UserInfoRes struct {
	UserId  int64  `json:"userId"`
	Name    string `json:"name"`
	Account string `json:"account"`
}
