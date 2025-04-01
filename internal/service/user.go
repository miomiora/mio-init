package service

import (
	"context"
	"fmt"
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

func (userService) Update(ctx context.Context, param *model.UserUpdateReq, userId int64) error {
	if userId != param.UserId {
		return fmt.Errorf("非法修改")
	}
	return repository.User.Update(ctx, &model.User{
		UserId: param.UserId,
		Name:   param.Name,
	})
}

func (userService) UpdatePwd(ctx context.Context, param *model.UserUpdatePwdReq, refreshToken, accessToken string) error {
	if param.NewPassword != param.RePassword {
		return fmt.Errorf("非法修改")
	}
	user, err := repository.User.Login(ctx, param.Account, util.Md5(param.Password))
	if err != nil {
		return err
	}
	user.Password = util.Md5(param.NewPassword)
	err = repository.User.Update(ctx, user)
	if err != nil {
		return err
	}

	return repository.Cache.Logout(ctx, refreshToken, accessToken)
}

func (userService) Delete(ctx context.Context, userId int64) error {
	return repository.User.Delete(ctx, userId)
}

func (userService) List(ctx context.Context, page, pageSize int, orderBy string) ([]*model.UserInfoRes, int64, error) {
	users, count, err := repository.User.GetAllUsers(ctx, page, pageSize, orderBy)
	if err != nil {
		return nil, 0, err
	}
	var result []*model.UserInfoRes
	for _, user := range users {
		result = append(result, &model.UserInfoRes{
			UserId:  user.UserId,
			Name:    user.Name,
			Account: user.Account,
		})
	}
	return result, count, nil
}
