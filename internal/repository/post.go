package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"mio-init/internal/model"
	"mio-init/internal/util"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) error

	GetByPostId(ctx context.Context, postId int64) (*model.Post, error)
	GetByTitle(ctx context.Context, title string, page, pageSize int) ([]*model.Post, int64, error)
	GetAllPosts(ctx context.Context, page, pageSize int, orderBy string) ([]*model.Post, int64, error)

	Update(ctx context.Context, post *model.Post) error
	Delete(ctx context.Context, postId int64) error
}

type postRepo struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepo{db: db}
}

func (r *postRepo) Create(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *postRepo) GetByPostId(ctx context.Context, postId int64) (*model.Post, error) {
	var post model.Post
	err := r.db.WithContext(ctx).Where("post_id = ?", postId).First(&model.Post{}).Error
	return &post, err
}

func (r *postRepo) GetByTitle(ctx context.Context, title string, page, pageSize int) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var count int64

	if err := r.db.WithContext(ctx).
		Model(&model.Post{}).
		Where("title = ?", title).
		Count(&count).Error; err != nil {
		return nil, 0, fmt.Errorf("count posts failed: %w", err)
	}

	offset := (page - 1) * pageSize
	if err := r.db.WithContext(ctx).
		Where("title = ?", title).
		Offset(offset).
		Limit(pageSize).
		Find(&posts).Error; err != nil {
		return nil, 0, fmt.Errorf("query posts failed: %w", err)
	}

	return posts, count, nil
}

func (r *postRepo) GetAllPosts(ctx context.Context, page, pageSize int, orderBy string) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var count int64

	// 1. 获取总数
	if err := r.db.WithContext(ctx).
		Model(&model.Post{}).
		Count(&count).Error; err != nil {
		return nil, 0, fmt.Errorf("count posts failed: %w", err)
	}

	// 2. 分页查询
	query := r.db.WithContext(ctx).Offset((page - 1) * pageSize).Limit(pageSize)

	// 3. 动态排序（安全校验）
	if orderBy != "" {
		if util.IsValidOrderField(orderBy) { // 防止SQL注入
			query = query.Order(orderBy)
		} else {
			query = query.Order("created_at DESC") // 默认排序
		}
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, 0, fmt.Errorf("query posts failed: %w", err)
	}

	return posts, count, nil
}

func (r *postRepo) Update(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Save(post).Error
}

func (r *postRepo) Delete(ctx context.Context, postId int64) error {
	return r.db.WithContext(ctx).Where("post_id = ?", postId).Delete(&model.Post{}).Error
}
