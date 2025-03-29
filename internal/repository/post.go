package repository

import (
	"context"
	"fmt"
	"mio-init/internal/core"
	"mio-init/internal/model"
	"mio-init/util"
)

type postRepo struct {
}

var Post = new(postRepo)

func (postRepo) Create(ctx context.Context, post *model.Post) error {
	return core.MySQL.GetDB().WithContext(ctx).Create(post).Error
}

func (postRepo) GetByPostId(ctx context.Context, postId int64) (*model.Post, error) {
	var post model.Post
	err := core.MySQL.GetDB().WithContext(ctx).Where("post_id = ?", postId).First(&model.Post{}).Error
	return &post, err
}

func (postRepo) GetByTitle(ctx context.Context, title string, page, pageSize int) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var count int64

	offset := (page - 1) * pageSize
	if err := core.MySQL.GetDB().WithContext(ctx).
		Where("title = ?", title).
		Offset(offset).
		Limit(pageSize).
		Find(&posts).Error; err != nil {
		return nil, 0, fmt.Errorf("query posts failed: %w", err)
	}

	return posts, count, nil
}

func (postRepo) GetAllPosts(ctx context.Context, page, pageSize int, orderBy string) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var count int64

	// 1. 获取总数
	if err := core.MySQL.GetDB().WithContext(ctx).
		Model(&model.Post{}).
		Count(&count).Error; err != nil {
		return nil, 0, fmt.Errorf("count posts failed: %w", err)
	}

	// 2. 分页查询
	query := core.MySQL.GetDB().WithContext(ctx).Offset((page - 1) * pageSize).Limit(pageSize)

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

func (postRepo) Update(ctx context.Context, post *model.Post) error {
	return core.MySQL.GetDB().WithContext(ctx).Save(post).Error
}

func (postRepo) Delete(ctx context.Context, postId int64) error {
	return core.MySQL.GetDB().WithContext(ctx).Where("post_id = ?", postId).Delete(&model.Post{}).Error
}
