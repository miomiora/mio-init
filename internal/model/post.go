package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	PostId   uint64
	Title    string
	Content  string
	AuthorId uint64
}
