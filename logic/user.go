package logic

import (
	"errors"
	"github.com/google/uuid"
	"mio-init/dao/mysql"
	"mio-init/dao/redis"
	"mio-init/model"
	"mio-init/util"
)

type userLogic struct {
}

var (
	ErrorUserExist = errors.New("用户已存在")
	User           = new(userLogic)
)

func (userLogic) Login(p *model.UserDTOLogin) (*model.UserVO, error) {
	// 1、从数据库中校验用户名和密码是否正确
	user, err := mysql.User.Login(p.Account, p.Password)
	if err != nil {
		return nil, err
	}

	// 2、登录成功，把 Token 存入 Redis 中
	token := uuid.NewString()
	err = redis.InsertTokenByUserId(token, user.UserId, user.UserRole)
	if err != nil {
		return nil, err
	}

	// 3、返回用户数据给 controller
	return &model.UserVO{
		UserId:      user.UserId,
		Account:     user.Account,
		Token:       &token,
		Description: nil,
		UserRole:    user.UserRole,
	}, nil
}

func (userLogic) Register(u *model.UserDTORegister) (err error) {
	// 判断用户存不存在
	if mysql.User.CheckAccountExist(u.Account) {
		return ErrorUserExist
	}
	// 生成userId
	userID := util.GenSnowflakeID()
	// 构造一个User实例
	user := &model.User{
		UserId:   userID,
		Account:  u.Account,
		Password: u.Password,
	}
	// 保存进数据库
	err = mysql.User.Insert(user)
	return
}

func (userLogic) Logout(token string) error {
	return redis.DeleteToken(token)
}

func (userLogic) GetUserList(params *model.ListParams) ([]*model.User, error) {
	return mysql.User.QueryUserList(params)
}

func (userLogic) GetLoginUser(userId int64) (*model.UserVO, error) {
	return mysql.User.QueryUserVOByUserId(userId)
}

func (userLogic) UpdateBySelf(u *model.UserDTOUpdateBySelf) error {
	user, err := mysql.User.QueryUserByUserId(u.UserId)
	if user.Account != u.Account {
		exist := mysql.User.CheckAccountExist(u.Account)
		if exist {
			return ErrorUserExist
		}
	}
	err = mysql.User.UpdateUserBySelf(u)
	return err
}

func (userLogic) GetUserVOByUserId(userId int64) (*model.UserVO, error) {
	return mysql.User.QueryUserVOByUserId(userId)
}

func (userLogic) GetUserVOList(params *model.ListParams) ([]*model.UserVO, error) {
	return mysql.User.QueryUserVOList(params)
}

func (userLogic) AddUser(u *model.UserDTOAdd) (err error) {
	// 与 register 一致。。。
	if mysql.User.CheckAccountExist(u.Account) {
		return ErrorUserExist
	}
	userID := util.GenSnowflakeID()
	user := &model.User{
		UserId:   userID,
		Account:  u.Account,
		Password: u.Password,
	}
	err = mysql.User.Insert(user)
	return
}

func (userLogic) DeleteUserByUserId(userId int64) error {
	return mysql.User.DeleteUserByUserId(userId)
}

func (userLogic) UpdateUserByAdmin(u *model.UserDTOUpdateByAdmin) error {
	return mysql.User.UpdateUserByAdmin(u)
}

func (userLogic) GetUserByUserId(userId int64) (*model.User, error) {
	return mysql.User.QueryUserByUserId(userId)
}
