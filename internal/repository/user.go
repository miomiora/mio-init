package repository

import (
	"context"
	"fmt"
	"mio-init/internal/core"
	"mio-init/internal/model"
	"mio-init/util"
)

type userRepo struct {
}

var User = new(userRepo)

func (userRepo) Create(ctx context.Context, user *model.User) error {
	return core.MySQL.GetDB().WithContext(ctx).Model(&model.User{}).Create(user).Error
}

func (userRepo) Login(ctx context.Context, account, password string) (*model.User, error) {
	var user model.User
	err := core.MySQL.GetDB().WithContext(ctx).Model(&model.User{}).Where("account = ? and password = ?", account, password).First(&user).Error
	return &user, err
}

func (userRepo) GetByUserId(ctx context.Context, userId int64) (*model.User, error) {
	var user model.User
	err := core.MySQL.GetDB().WithContext(ctx).Model(&model.User{}).Where("user_id = ?", userId).First(&user).Error
	return &user, err
}

func (userRepo) GetByName(ctx context.Context, name string, page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	var count int64

	offset := (page - 1) * pageSize
	if err := core.MySQL.GetDB().WithContext(ctx).Model(&model.User{}).
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
	if err := core.MySQL.GetDB().WithContext(ctx).Model(&model.User{}).
		Model(&model.User{}).
		Count(&count).Error; err != nil {
		return nil, 0, fmt.Errorf("count users failed: %w", err)
	}

	// 2. 分页查询
	query := core.MySQL.GetDB().WithContext(ctx).Model(&model.User{}).Offset((page - 1) * pageSize).Limit(pageSize)

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
	return core.MySQL.GetDB().WithContext(ctx).Model(&model.User{}).Where("user_id = ?", user.UserId).Updates(user).Error
}

func (userRepo) Delete(ctx context.Context, userId int64) error {
	return core.MySQL.GetDB().WithContext(ctx).Where("user_id = ?", userId).Delete(&model.User{}).Error
}
