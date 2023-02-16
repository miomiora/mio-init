package controllers

import (
	"gorm.io/gorm"
	"user-center/models"
)

type Response struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	IsSuccess bool        `json:"is_success"`
}

var db *gorm.DB

func init() {
	db = models.DB
}
