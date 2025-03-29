package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mio-init/model"
	"mio-init/service"
	"mio-init/util"
	"strconv"
)

type postController struct {
}

var Post = new(postController)

// InsertPost
// @Summary 新建文章
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param object body model.PostDTOInsert true "新建文章参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/new [post]
func (postController) InsertPost(c *gin.Context) {
	// 1、参数校验
	p := new(model.PostDTOInsert)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误
		zap.L().Warn("[controller postController InsertPost] insert post with invalid param ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	// 获取userId
	userId, err := getUserId(c)
	if err != nil {
		zap.L().Warn("[controller postController InsertPost] get userId error ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	// 2、业务处理
	if err = service.Post.InsertPost(p, userId); err != nil {
		zap.L().Warn("[controller postController InsertPost] insert post failed ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	// 3、返回响应
	ResponseOK(c, nil)
}

// UpdateBySelf
// @Summary 用户更新自己写的文章
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param object body model.PostDTOUpdateBySelf true "修改后的数据"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/update/my [post]
func (postController) UpdateBySelf(c *gin.Context) {
	// 验证参数
	u := new(model.PostDTOUpdateBySelf)
	err := c.ShouldBindJSON(u)
	if err != nil {
		zap.L().Warn("[controller postController UpdateBySelf] update post by self with invalid param ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	// 验证是否本人
	userId, err := getUserId(c)
	if err != nil || userId != u.UserId {
		zap.L().Warn("[controller postController UpdateBySelf] get userId error ", zap.Error(err))
		ResponseError(c, ErrorNotLogin)
		return
	}
	// 业务
	if err = service.Post.UpdateBySelf(u); err != nil {
		zap.L().Warn("[controller postController UpdateBySelf] update post by self failed ", zap.Error(err))
		if errors.Is(err, service.ErrorPostExist) {
			ResponseErrorWithMsg(c, ErrorInvalidParams, err.Error())
			return
		}
	}
	ResponseOK(c, nil)
}

// GetPostVOByPostId
// @Summary 通过postId获取文章视图
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param postId query string true "需要查找的文章id"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/get/vo [get]
func (postController) GetPostVOByPostId(c *gin.Context) {
	value := c.Query(util.KeyPostId)
	if value == "" {
		zap.L().Warn("[controller postController GetPostVOByPostId] query postId failed ")
		ResponseError(c, ErrorInvalidParams)
		return
	}
	postId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		zap.L().Warn("[controller postController GetPostVOByPostId] parse postId failed ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	data, err := service.Post.GetPostVOByPostId(postId)
	if err != nil {
		zap.L().Warn("[controller postController GetPostVOByPostId] get post vo by postId failed ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	ResponseOK(c, data)
}

// GetMyPostVOList
// @Summary 通过当前登录的用户所写的文章
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param object body model.ListParams true "分页查询需要的参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/my [post]
func (postController) GetMyPostVOList(c *gin.Context) {
	params := new(model.ListParams)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Warn("[controller postController GetMyPostVOList] get my post vo list with invalid param ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		zap.L().Warn("[controller postController GetMyPostVOList] get userId error ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	data, err := service.Post.GetMyPostVOList(params, userId)
	if err != nil {
		zap.L().Warn("[controller postController GetMyPostVOList] get my post vo list failed ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	ResponseOK(c, data)
}

// GetPostVOList
// @Summary 通过获取文章视图列表
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生"
// @Param object body model.ListParams true "分页查询需要的参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/list/page/vo [post]
func (postController) GetPostVOList(c *gin.Context) {
	params := new(model.ListParams)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Warn("[controller postController GetPostVOList] get post vo list with invalid param ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	data, err := service.Post.GetPostVOList(params)
	if err != nil {
		zap.L().Warn("[controller postController GetPostVOList] get post vo list failed ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	ResponseOK(c, data)
}

// AddPost
// @Summary 通过获取文章视图列表
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param object body model.PostDTOAdd true "新增的文章参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/add [post]
func (postController) AddPost(c *gin.Context) {
	// 1、参数校验
	u := new(model.PostDTOAdd)
	if err := c.ShouldBindJSON(u); err != nil {
		// 请求参数有误
		zap.L().Error("[controller post AddPost] add post with invalid param ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		zap.L().Warn("[controller post AddPost] get userId error ")
		ResponseError(c, ErrorServerBusy)
		return
	}
	// 2、业务处理
	if err = service.Post.AddPost(u, userId); err != nil {
		zap.L().Error("[controller post AddPost] add post failed ", zap.Error(err))
		if errors.Is(err, service.ErrorPostExist) {
			ResponseErrorWithMsg(c, ErrorInvalidParams, err.Error())
			return
		}
		ResponseError(c, ErrorServerBusy)
		return
	}
	// 3、返回响应
	ResponseOK(c, nil)
}

// DeletePostByPostId
// @Summary 通过postId删除文章
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param postId query string true "需要删除的postId"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/delete [post]
func (postController) DeletePostByPostId(c *gin.Context) {
	value := c.Query(util.KeyPostId)
	if value == "" {
		zap.L().Warn("[controller postController DeletePostByPostId] query postId failed ")
		ResponseError(c, ErrorInvalidParams)
		return
	}
	postId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		zap.L().Warn("[controller postController DeletePostByPostId] parse postId failed ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	if err = service.Post.DeletePostByPostId(postId); err != nil {
		zap.L().Warn("[controller postController DeletePostByPostId] delete post by postId failed ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	ResponseOK(c, nil)
}

// UpdatePostByAdmin
// @Summary 管理员编辑文章
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param object body model.PostDTOUpdateByAdmin true "需要更新的文章信息"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/update [post]
func (postController) UpdatePostByAdmin(c *gin.Context) {
	u := new(model.PostDTOUpdateByAdmin)
	err := c.ShouldBindJSON(u)
	if err != nil {
		zap.L().Warn("[controller postController UpdatePostByAdmin] update post by admin with invalid param ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	// 业务
	err = service.Post.UpdatePostByAdmin(u)
	if err != nil {
		zap.L().Warn("[controller postController UpdatePostByAdmin] update post by admin failed ", zap.Error(err))
		if errors.Is(err, service.ErrorPostExist) {
			ResponseErrorWithMsg(c, ErrorInvalidParams, err.Error())
			return
		}
	}
	ResponseOK(c, nil)
}

// GetPostByPostId
// @Summary 管理员通过postId获取文章全部数据
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param postId query string true "需要查询的postId"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/get [get]
func (postController) GetPostByPostId(c *gin.Context) {
	value := c.Query(util.KeyPostId)
	if value == "" {
		zap.L().Warn("[controller postController GetPostByPostId] query postId failed ")
		ResponseError(c, ErrorInvalidParams)
		return
	}
	postId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		zap.L().Warn("[controller postController GetPostByPostId] parse postId failed ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	data, err := service.Post.GetPostByPostId(postId)
	if err != nil {
		zap.L().Warn("[controller postController GetPostByPostId] get post by postId failed ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	ResponseOK(c, data)
}

// GetPostList
// @Summary 管理员获取全部文章详细信息
// @Tags 文章相关接口
// @Accept json
// @Produce json
// @Param Authorization header string true "用户令牌 Token 登录后产生，需为管理员"
// @Param object body model.ListParams true "分页查询所需要的参数"
// @Security ApiKeyAuth
// @Success 200 {object} Response
// @Router /post/list/page [post]
func (postController) GetPostList(c *gin.Context) {
	params := new(model.ListParams)
	err := c.ShouldBindJSON(params)
	if err != nil {
		zap.L().Warn("[controller postController GetPostList] get post list with invalid param ", zap.Error(err))
		ResponseError(c, ErrorInvalidParams)
		return
	}
	data, err := service.Post.GetPostList(params)
	if err != nil {
		zap.L().Warn("[controller postController GetPostList] get post list failed ", zap.Error(err))
		ResponseError(c, ErrorServerBusy)
		return
	}
	ResponseOK(c, data)
}
