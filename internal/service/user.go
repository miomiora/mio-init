package service

import (
	"context"
	"mio-init/internal/model"
	"mio-init/internal/repository"
	"mio-init/util"
	"strconv"
)

type userService struct {
}

var User = new(userService)

func (userService) Login(ctx context.Context, req *model.UserLoginReq) (*model.UserLoginRes, error) {
	user, err := repository.User.Login(ctx, req.Account, util.Md5(req.Password))
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := util.GenTokens(user.UserId)
	if err != nil {
		return nil, err
	}

	_, err = repository.Cache.Set(ctx, util.GenRefreshKey(refreshToken),
		strconv.FormatInt(user.UserId, 10), util.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}

	return &model.UserLoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       user.UserId,
		Name:         user.Name,
		Account:      user.Account,
	}, nil
}

func (userService) Create(ctx context.Context, req *model.UserCreateReq) error {
	return repository.User.Create(ctx, &model.User{
		UserId:   util.GenSnowflakeID(),
		Name:     req.Name,
		Account:  req.Account,
		Password: util.Md5(req.Password),
	})
}

func (userService) Logout(ctx context.Context, refreshToken, accessToken string) error {
	return repository.Cache.Logout(ctx, refreshToken, accessToken)
}

func (userService) GetByUserId(ctx context.Context, userId int64) (*model.UserInfoRes, error) {
	user, err := repository.User.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &model.UserInfoRes{
		UserId:  user.UserId,
		Name:    user.Name,
		Account: user.Account,
	}, nil
}
