package repository

import "gorm.io/gorm"

type Repositories struct {
	Post PostRepository
	User PostRepository
	// 其他 Repository...
}

// NewRepositories 集中初始化所有 Repository
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Post: NewPostRepository(db),
		User: NewPostRepository(db),
	}
}
