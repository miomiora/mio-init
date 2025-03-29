package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserId   uint64
	Name     string
	Email    string
	Account  string
	Password string
}
