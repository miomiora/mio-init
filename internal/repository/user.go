package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"mio-init/internal/model"
	"mio-init/internal/util"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error

	Login(ctx context.Context, account, password string) (*model.User, error)

	GetByUserId(ctx context.Context, userId int64) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByName(ctx context.Context, name string, page, pageSize int) ([]*model.User, int64, error)
	GetAllUsers(ctx context.Context, page, pageSize int, orderBy string) ([]*model.User, int64, error)

	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, userId int64) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) Login(ctx context.Context, account, password string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("account = ? and password = ?", account, password).First(&model.User{}).Error
	return &user, err
}

func (r *userRepo) GetByUserId(ctx context.Context, userId int64) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&model.User{}).Error
	return &user, err
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepo) GetByName(ctx context.Context, name string, page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	var count int64

	if err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("name = ?", name).
		Count(&count).Error; err != nil {
		return nil, 0, fmt.Errorf("count users failed: %w", err)
	}

	offset := (page - 1) * pageSize
	if err := r.db.WithContext(ctx).
		Where("name = ?", name).
		Offset(offset).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("query users failed: %w", err)
	}

	return users, count, nil
}

func (r *userRepo) GetAllUsers(ctx context.Context, page, pageSize int, orderBy string) ([]*model.User, int64, error) {
	var users []*model.User
	var count int64

	// 1. 获取总数
	if err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Count(&count).Error; err != nil {
		return nil, 0, fmt.Errorf("count users failed: %w", err)
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

	if err := query.Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("query users failed: %w", err)
	}

	return users, count, nil
}

func (r *userRepo) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepo) Delete(ctx context.Context, userId int64) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userId).Delete(&model.User{}).Error
}
