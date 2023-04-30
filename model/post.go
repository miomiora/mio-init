package model

import (
	"gorm.io/gorm"
	"time"
)

// Post 文章
type Post struct {
	gorm.Model

	PostId  int64  `json:"post_id" gorm:"not null"`
	UserId  int64  `json:"user_id" gorm:"not null"`
	Title   string `json:"title" gorm:"not null;size:128"`
	Content string `json:"content" gorm:"not null"`
}

// PostDTOInsert 创建新的文章
type PostDTOInsert struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostDTOAdd struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostDTOUpdateBySelf struct {
	UserId  int64  `json:"user_id,string" binding:"required"`
	PostId  int64  `json:"post_id,string" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostDTOUpdateByAdmin struct {
	PostId  int64  `json:"post_id,string" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// PostVO 返回给前端响应的文章信息
type PostVO struct {
	Account   string    `json:"account"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	PostId    int64     `json:"post_id,string"`
	CreatedAt time.Time `json:"created_at"`
}
