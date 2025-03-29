package service

import (
	"context"
	"mio-init/internal/model"
	"mio-init/internal/repository"
	"mio-init/util"
)

type userService struct {
}

var User = new(userService)

func (userService) Login(ctx context.Context, req *model.UserLoginReq) (*model.User, error) {
	// 1、从数据库中校验用户名和密码是否正确
	user, err := repository.User.Login(ctx, req.Account, util.EncryptStr(req.Password))
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := util.GenerateTokens(user.UserId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService) Create(ctx context.Context, req *model.UserCreateReq) error {
	return repository.User.Create(ctx, &model.User{
		UserId:   util.GenSnowflakeID(),
		Name:     req.Name,
		Account:  req.Account,
		Password: util.EncryptStr(req.Password),
	})
}

func (userService) Logout(ctx context.Context, userId int64) error {
	return nil
}
