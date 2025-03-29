package service

import (
	"errors"
	"mio-init/dao/mysql"
	"mio-init/model"
	"mio-init/util"
)

type postLogic struct {
}

var (
	ErrorPostExist = errors.New("文章已存在")
	Post           = new(postLogic)
)

func (postLogic) InsertPost(p *model.PostDTOInsert, userId int64) (err error) {
	postId := util.GenSnowflakeID()
	// 构造一个User实例
	post := &model.Post{
		UserId:  userId,
		PostId:  postId,
		Content: p.Content,
		Title:   p.Title,
	}
	// 保存进数据库
	err = mysql.Post.Insert(post)
	return
}

func (postLogic) GetPostVOList(params *model.ListParams) ([]*model.PostVO, error) {
	data, err := mysql.Post.QueryPostList(params)
	if err != nil {
		return nil, err
	}

	var postList []*model.PostVO
	var user *model.User

	for _, value := range data {
		user, err = mysql.User.QueryUserByUserId(value.UserId)
		if err != nil {
			return nil, err
		}

		post := &model.PostVO{
			Account:   user.Account,
			Title:     value.Title,
			Content:   value.Content,
			PostId:    value.PostId,
			CreatedAt: value.CreatedAt,
		}

		postList = append(postList, post)
	}
	return postList, nil
}

func (postLogic) GetLoginPost(postId int64) (*model.PostVO, error) {
	return mysql.Post.QueryPostVOByPostId(postId)
}

func (postLogic) UpdateBySelf(u *model.PostDTOUpdateBySelf) error {
	return mysql.Post.UpdatePostBySelf(u)
}

func (postLogic) GetPostVOByPostId(postId int64) (*model.PostVO, error) {
	return mysql.Post.QueryPostVOByPostId(postId)
}

func (postLogic) GetMyPostVOList(params *model.ListParams, userId int64) ([]*model.PostVO, error) {
	return mysql.Post.QueryPostVOListByUserId(params, userId)
}

func (postLogic) GetPostList(params *model.ListParams) ([]*model.Post, error) {
	return mysql.Post.QueryPostList(params)
}

func (postLogic) AddPost(u *model.PostDTOAdd, userId int64) (err error) {
	// 与 register 一致。。。
	postID := util.GenSnowflakeID()
	post := &model.Post{
		PostId:  postID,
		Content: u.Content,
		Title:   u.Title,
		UserId:  userId,
	}
	err = mysql.Post.Insert(post)
	return
}

func (postLogic) DeletePostByPostId(postId int64) error {
	return mysql.Post.DeletePostByPostId(postId)
}

func (postLogic) UpdatePostByAdmin(u *model.PostDTOUpdateByAdmin) error {
	return mysql.Post.UpdatePostByAdmin(u)
}

func (postLogic) GetPostByPostId(postId int64) (*model.Post, error) {
	return mysql.Post.QueryPostByPostId(postId)
}
