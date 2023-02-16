package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"regexp"
	"user-center/models"
)

func UserLogin(c *gin.Context) {
	var userDTO models.UserDTO
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, Response{
			Message:   "接收User对象失败！" + err.Error(),
			Data:      nil,
			IsSuccess: false,
		})
		return
	}
	// 查询数据库中是否存在该用户，并且同时把取出来的数据存入userDTO中
	affected := db.
		Take(&models.User{}, "name = ? and password = ?", user.UserName, user.UserPassword).
		Scan(&userDTO).RowsAffected
	if affected == 0 {
		c.JSON(500, Response{
			Message:   "用户名不存在或密码错误！",
			Data:      nil,
			IsSuccess: false,
		})
		return
	}

	//token, err := middleware.GetToken(user)
	//if err != nil {
	//	c.JSON(500, Response{
	//		Message:   err.Error(),
	//		Data:      nil,
	//		IsSuccess: false,
	//	})
	//	return
	//}

	c.JSON(200, Response{
		Message:   "用户登录成功",
		Data:      userDTO,
		IsSuccess: true,
	})
}

func UserRegister(c *gin.Context) {
	var userRegister models.UserRegister
	// 获取用户注册信息, 同时校验是否为空, 以及长度是否合法
	err := c.ShouldBindJSON(&userRegister)
	if err != nil {
		c.JSON(500, gin.H{"message": "获取用户注册信息失败！" + err.Error()})
		return
	}

	// 帐号是否合法(字母开头，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]*$
	matched, err := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, userRegister.UserAccount)
	if !matched {
		c.JSON(500, gin.H{"message": "账号不合法！"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"message": "校验账号失败！" + err.Error()})
		return
	}
	// 密码(以字母开头，只能包含字母、数字和下划线)：^[a-zA-Z]\w*$    \w = [a-zA-Z0-9_]

	// 账户是否重复
	affected := db.Take(&models.User{}, "user_account = ?", userRegister.UserAccount).RowsAffected
	if affected == 1 {
		c.JSON(500, gin.H{"message": "用户已经存在！"})
		return
	}

	// 密码加密
	hash := md5.New()
	hash.Write([]byte("miomio"))                  // 先加盐
	hash.Write([]byte(userRegister.UserPassword)) // 再加密密码
	encryptPassword := hex.EncodeToString(hash.Sum(nil))

	// 插入数据
	user := &models.User{
		UserPassword: encryptPassword,
		UserAccount:  userRegister.UserAccount,
	}
	affected = db.Save(user).RowsAffected
	if affected == 1 {
		c.JSON(200, gin.H{
			"message": "注册成功",
			"data":    user,
		})
	}

}
