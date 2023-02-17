package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"net/http"
	"regexp"
	"user-center/models"
	"user-center/utils"
)

var salt = "miomio"

func UserLogin(c *gin.Context) {
	var userDTO models.UserDTO
	var user models.User
	// 获取用户登录信息, 同时校验是否为空, 以及长度是否合法
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "获取用户登录信息失败！" + err.Error(),
		})
		return
	}

	// 帐号是否合法(字母开头，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]*$
	matched, err := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, user.UserAccount)
	if !matched {
		c.JSON(http.StatusForbidden, gin.H{"message": "账号不合法！"})
		return
	}
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "校验账号失败！" + err.Error()})
		return
	}

	hash := md5.New()
	hash.Write([]byte(salt))              // 先加盐
	hash.Write([]byte(user.UserPassword)) // 再加密密码
	encryptPassword := hex.EncodeToString(hash.Sum(nil))
	// 查询数据库中是否存在该用户，并且同时把取出来的数据存入userDTO中
	affected := DB.
		Take(&models.User{},
			"user_account = ? and user_password = ?", user.UserAccount, encryptPassword).
		Scan(&userDTO).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "用户名不存在或密码错误",
		})
		return
	}

	// 记录用户的登录状态, 使用redis+token
	token := uuid.NewString()
	tokenKey := utils.TOKEN_PREIX + token
	// 存入redis, 并且把用户ip存入redis
	_, err = Conn.Do("HSET", tokenKey,
		"id", userDTO.ID,
		"client_ip", c.ClientIP())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "存储Token失败！" + err.Error(),
			"data":    nil,
		})
		return
	}
	// 设置有效期
	_, err = Conn.Do("EXPIRE", tokenKey, utils.TOKEN_TIMEOUT)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "设置Token有效期失败！" + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"data":    userDTO,
	})
}

func UserRegister(c *gin.Context) {
	var userRegister models.UserRegister
	// 获取用户注册信息, 同时校验是否为空, 以及长度是否合法
	err := c.ShouldBindJSON(&userRegister)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "获取用户注册信息失败！" + err.Error()})
		return
	}

	// 帐号是否合法(字母开头，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]*$
	matched, err := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, userRegister.UserAccount)
	if !matched {
		c.JSON(http.StatusForbidden, gin.H{"message": "账号不合法！"})
		return
	}
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "校验账号失败！" + err.Error()})
		return
	}
	// 密码(以字母开头，只能包含字母、数字和下划线)：^[a-zA-Z]\w*$    \w = [a-zA-Z0-9_]

	// 账户是否重复
	affected := DB.Take(&models.User{}, "user_account = ?", userRegister.UserAccount).RowsAffected
	if affected == 1 {
		c.JSON(http.StatusForbidden, gin.H{"message": "用户已经存在！"})
		return
	}

	// 密码加密
	hash := md5.New()
	hash.Write([]byte(salt))                      // 先加盐
	hash.Write([]byte(userRegister.UserPassword)) // 再加密密码
	encryptPassword := hex.EncodeToString(hash.Sum(nil))

	// 插入数据
	user := &models.User{
		UserPassword: encryptPassword,
		UserAccount:  userRegister.UserAccount,
	}
	affected = DB.Save(user).RowsAffected
	if affected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
			"data":    user,
		})
	}
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var userDTO models.UserDTO
	affected := DB.Take(&models.User{}, "id = ?", id).Scan(&userDTO).RowsAffected
	if affected == 0 {
		c.JSON(500, gin.H{"message": "用户不存在！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("查询用户%v成功", id),
		"data":    userDTO,
	})
}

func GetUserList(c *gin.Context) {
	var userList []models.UserDTO
	affected := DB.Model(&models.User{}).Scan(&userList).RowsAffected
	if affected == 0 {
		c.JSON(500, gin.H{"message": "查询失败！"})
		return
	}
	if affected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "查询成功！",
			"users":   userList,
		})
		return
	}
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	// 判断是否删除的用户是自己
	token := c.GetHeader("Authorization")
	tokenKey := utils.TOKEN_PREIX + token
	// 获取发起请求的用户id
	userId, err := redis.String(Conn.Do("HGET", tokenKey, "id"))
	if err != nil {
		fmt.Println("DeleteUserById redis.String(Conn.Do(\"HMGET\", tokenKey, \"id\")) err!!!" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "验证删除用户权限失败！"})
		return
	}
	// 如果删除的id和发起请求人的id一致则不能删除
	if userId == id {
		c.JSON(http.StatusForbidden, gin.H{"message": "不能删除自己！"})
		return
	}
	// 删除用户
	affected := DB.Delete(&models.User{}, "id = ?", id).RowsAffected
	// 用户已经不存在了
	if affected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func GetUserLoginPage(c *gin.Context) {
	// 能走到这个函数，说明已经验证过用户存在，直接给前端返回200
	c.JSON(http.StatusOK, gin.H{"message": "用户已经登录！"})
}
