package repository

import (
	"context"
	"fmt"
	"mio-init/internal/core"
	"mio-init/internal/model"
	"mio-init/internal/util"
)

type userRepo struct {
}

var UserRepo = new(userRepo)

func (userRepo) Create(ctx context.Context, user *model.User) error {
	return core.GetDB().WithContext(ctx).Create(user).Error
}

func (userRepo) Login(ctx context.Context, account, password string) (*model.User, error) {
	var user model.User
	err := core.GetDB().WithContext(ctx).Where("account = ? and password = ?", account, password).First(&model.User{}).Error
	return &user, err
}

func (userRepo) GetByUserId(ctx context.Context, userId int64) (*model.User, error) {
	var user model.User
	err := core.GetDB().WithContext(ctx).Where("user_id = ?", userId).First(&model.User{}).Error
	return &user, err
}

func (userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := core.GetDB().WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (userRepo) GetByName(ctx context.Context, name string, page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	var count int64

	offset := (page - 1) * pageSize
	if err := core.GetDB().WithContext(ctx).
		Where("name = ?", name).
		Offset(offset).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("query users failed: %w", err)
	}

	return users, count, nil
}

func (userRepo) GetAllUsers(ctx context.Context, page, pageSize int, orderBy string) ([]*model.User, int64, error) {
	var users []*model.User
	var count int64

	// 1. 获取总数
	if err := core.GetDB().WithContext(ctx).
		Model(&model.User{}).
		Count(&count).Error; err != nil {
		return nil, 0, fmt.Errorf("count users failed: %w", err)
	}

	// 2. 分页查询
	query := core.GetDB().WithContext(ctx).Offset((page - 1) * pageSize).Limit(pageSize)

	// 3. 动态排序（安全校验）
	if orderBy != "" {
		if util.IsValidOrderField(orderBy) { // 防止SQL注入
			query = query.Order(orderBy)
		} else {
			query = query.Order("created_at DESC") // 默认排序
		}
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("query users failed: %w", err)
	}

	return users, count, nil
}

func (userRepo) Update(ctx context.Context, user *model.User) error {
	return core.GetDB().WithContext(ctx).Save(user).Error
}

func (userRepo) Delete(ctx context.Context, userId int64) error {
	return core.GetDB().WithContext(ctx).Where("user_id = ?", userId).Delete(&model.User{}).Error
}
