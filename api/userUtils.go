package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-center/models"
	"user-center/utils"
)

/*
	主要对userApi中常出现的内容进行封装
*/

//
// isUserAccountExist
//  @Description: 判断用户名是否存在
//  @param c
//  @param userAccount 判断是否存在的用户名
//  @return bool 返回用户是否存在，true为存在，false表示未存
//
func isUserAccountExist(c *gin.Context, userAccount string) bool {
	affected := DB.Take(&models.User{}, "user_account = ?", userAccount).RowsAffected
	if affected == 1 {
		c.JSON(http.StatusForbidden,
			utils.ResponseError(utils.ParamsError, "账号已存在！"))
		return true
	}
	return false
}

//
// encryptPassword
//  @Description: 密码加密
//  @param password 需要加密的密码
//  @return string 加密后的密码
//
func encryptPassword(password string) string {
	// 密码加盐
	var salt = "miomio"
	hash := md5.New()
	hash.Write([]byte(salt))     // 先加盐
	hash.Write([]byte(password)) // 再加密密码
	encryptString := hex.EncodeToString(hash.Sum(nil))
	return encryptString
}

//
// getContextValue
//  @Description: 用于获取中间件设置在context中的值
//  @param c
//  @param key 需要获取的值的键
//  @return interface{} 获取到的值，返回nil则表示未获取到值
//
func getContextValue(c *gin.Context, key string) interface{} {
	value, exists := c.Get(key)
	if !exists {
		c.JSON(http.StatusInternalServerError,
			utils.ResponseError(utils.ServerError, "获取用户身份失败！"))
		return nil
	}
	return value
}

//
// matchParamId
//  @Description: 绑定请求传来的id参数
//  @param c
//  @return string 返回空字符串表示匹配参数失败
//
func matchParamId(c *gin.Context) string {
	id := c.Param("id")
	matched := utils.MatchString(`^[0-9]*$`, id)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "id必须为数字！"))
		return ""
	}
	return id
}

//
// bindContextJson
//  @Description: 用来绑定context的json值，并做相关的错误处理
//  @param c
//  @param data 需要把context的json值绑定到什么变量上
//  @return bool 返回true为成功，false为失败
//
func bindContextJson(c *gin.Context, data interface{}) bool {
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println("[api bindContextJson err] c.ShouldBindJSON : ", err.Error())
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "缺少必填参数或不合法！"))
		return false
	}
	return true
}

//
// updateUserInfo
//  @Description: 更新用户的基本信息
//  @param c
//  @param user 需要更新后的值
//  @return bool 返回true表示成功，false表示失败
//
func updateUserInfo(c *gin.Context, user models.UserDTO) bool {
	//零值不更新
	affected := DB.Take(&models.User{}, user.ID).Updates(models.User{
		UserAccount: user.UserAccount,
		UserName:    user.UserName,
		AvatarUrl:   user.AvatarUrl,
		Gender:      user.Gender,
		Phone:       user.Phone,
		Email:       user.Email,
		UserStatus:  user.UserStatus,
		Role:        user.Role,
	}).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "Mysql修改用户信息错误！"))
		return false
	}
	return true
}

func changePassword(c *gin.Context, user models.UserChangePassword) bool {
	// 判断密码和确认密码是否一致
	if user.UserPassword != user.CheckPassword {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "两次密码不一样！"))
		return false
	}
	// 密码加密
	password := encryptPassword(user.UserPassword)
	// 修改
	affected := DB.Take(&models.User{}, user.ID).Updates(models.User{
		UserPassword: password,
	}).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "Mysql修改用户信息错误！"))
		return false
	}
	return true
}
