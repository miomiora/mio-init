package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"user-center/models"
	"user-center/utils"
)

//
// UserLogin
//  @Description: 用户登录, 需要接受json内容:user_account、user_password
//  @param c
//
func UserLogin(c *gin.Context) {
	var userDTO models.UserDTO
	var user models.User
	// 获取用户登录信息, 同时校验是否为空, 以及长度是否合法
	result := bindContextJson(c, &user)
	if !result {
		return
	}

	// 帐号是否合法(字母开头，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]*$
	matched := utils.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, user.UserAccount)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不合法！"))
		return
	}
	// 加密密码
	password := encryptPassword(user.UserPassword)

	// 查询数据库中是否存在该用户，并且同时把取出来的数据存入userDTO中
	affected := DB.
		Take(&models.User{},
			"user_account = ? and user_password = ?", user.UserAccount, password).
		Scan(&userDTO).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不存在！"))
		return
	}

	// 记录用户的登录状态, 使用redis+token
	token := uuid.NewString()
	tokenKey := utils.TokenPrefix + token
	// 存入redis, 并且把用户ip存入redis
	_, err := Conn.Do("HSET", tokenKey, "id", userDTO.ID, "client_ip", c.ClientIP())

	if err != nil {
		fmt.Println("[api UserLogin err] Conn.Do HSET : " + err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.RedisError, "存储Token失败！"))
		return
	}
	// 设置有效期
	_, err = Conn.Do("EXPIRE", tokenKey, utils.TokenTimeout)
	if err != nil {
		fmt.Println("[api UserLogin err] Conn.Do EXPIRE : " + err.Error())
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.RedisError, "设置Token有效期失败！"))
		return
	}

	// 封装user和token
	res := &gin.H{
		"user":  userDTO,
		"token": token,
	}
	c.JSON(http.StatusOK, utils.ResponseOK(res))
}

//
// UserRegister
//  @Description: 用户注册, 需要接受json内容:user_account、user_password、check_password
//  @param c
//
func UserRegister(c *gin.Context) {
	var userRegister models.UserRegister
	// 获取用户注册信息, 同时校验是否为空, 以及长度是否合法
	result := bindContextJson(c, &userRegister)
	if !result {
		return
	}
	// 帐号是否合法(字母开头，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]*$
	matched := utils.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, userRegister.UserAccount)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "账号不合法！"))
		return
	}
	// 密码(以字母开头，只能包含字母、数字和下划线)：^[a-zA-Z]\w*$    \w = [a-zA-Z0-9_]

	// 账户是否重复
	exist := isUserAccountExist(c, userRegister.UserAccount)
	if exist {
		return
	}

	// 密码加密
	password := encryptPassword(userRegister.UserPassword)

	var userDTO models.UserDTO
	// 插入数据
	user := &models.User{
		UserPassword: password,
		UserAccount:  userRegister.UserAccount,
	}
	affected := DB.Save(user).Scan(&userDTO).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "注册用户失败！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(userDTO))
}

//
// GetUserById
//  @Description: 获取指定id的用户, 需要接受请求参数:id
//  @param c
//
func GetUserById(c *gin.Context) {
	// 验证id是否合法
	id := matchParamId(c)
	if id == "" {
		return
	}

	var userDTO models.UserDTO
	affected := DB.Take(&models.User{}, id).Scan(&userDTO).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "该用户不存在！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(userDTO))
}

//
// GetUserList
//  @Description: 获取全部用户，需要接收请求参数: num一页的数量，page当前的页数
//  @param c
//
func GetUserList(c *gin.Context) {
	// 获取num和page的参数，并验证
	numParam := c.Param("num")
	matched := utils.MatchString(`^[0-9]*$`, numParam)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "num必须为数字！"))
		return
	}
	pageParam := c.Param("page")
	matched = utils.MatchString(`^[0-9]*$`, pageParam)
	if !matched {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "page必须为数字！"))
		return
	}

	// 将num和page转化为整数
	num, err := strconv.Atoi(numParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "转化为数字失败！"))
		return
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "转化为数字失败！"))
		return
	}

	var userList []models.UserDTO
	offset := (page - 1) * num
	affected := DB.Limit(num).Offset(offset).Model(&models.User{}).Scan(&userList).RowsAffected
	fmt.Println(affected)
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "数据库中没有用户！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(userList, fmt.Sprintf("查找到了%v个用户！", affected)))
}

//
// DeleteUserById
//  @Description: 删除指定用户，需要接受请求参数: id
//  @param c
//
func DeleteUserById(c *gin.Context) {
	//验证id是否合法
	id := matchParamId(c)
	if id == "" {
		return
	}
	// 获取发起请求的用户id
	userId := getContextValue(c, "userId")
	if userId == nil {
		return
	}
	// 判断是否删除的用户是自己 如果删除的id和发起请求人的id一致则不能删除
	if userId == id {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "不能删除自己！"))
		return
	}
	// 删除用户
	affected := DB.Delete(&models.User{}, id).RowsAffected
	// 用户已经不存在了
	if affected == 0 {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "该用户已经不存在！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil))
}

//
// GetCurrentUser
//  @Description: 获取当前用户,无需前端传参数,将从token中取得当前的用户
//  @param c
//
func GetCurrentUser(c *gin.Context) {
	userId, exists := c.Get("userId")
	var user models.UserDTO
	if !exists {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.ServerError, "获取用户身份失败！"))
		return
	}
	affected := DB.Take(&models.User{}, userId).Scan(&user).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "该用户不存在！"))
		return
	}
	// 能走到这个函数，说明已经验证过用户存在，直接给前端返回200
	c.JSON(http.StatusOK, utils.ResponseOK(user))
}

//
// UserLogout
//  @Description: 用户登出,无需前端传参数,将从token中取得当前的用户
//  @param c
//
func UserLogout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	tokenKey := utils.TokenPrefix + token
	if token != "" {
		_, err := Conn.Do("HDEL", tokenKey, "client_ip", "id")
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.RedisError, "用户Token已失效！"))
		}
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil))
}

//
// UpdateUserBySelf
//  @Description: 用户修改自己的信息,需验证token,
// 				  并且需要JSON参数: 	id(必须),user_account(必须),user_name,gender,phone,email,avatar_url
//  @param c
//
func UpdateUserBySelf(c *gin.Context) {
	// 获取当前发起请求的用户id
	userId := getContextValue(c, "userId")
	if userId == nil {
		return
	}
	// 获取当前发起请求的用户账号
	userAccount := getContextValue(c, "userAccount")
	if userAccount == nil {
		return
	}

	// 获取前端发送的用户信息
	var user models.UserDTO
	result := bindContextJson(c, &user)
	if !result {
		return
	}

	// 判断是否是要修改的用户本人发起的请求
	if userId != strconv.Itoa(int(user.ID)) {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "身份验证失败！"))
		return
	}

	// 用户是否修改了账户名，需要校验账户名是否存在
	if user.UserAccount != userAccount {
		exist := isUserAccountExist(c, user.UserAccount)
		if exist {
			return
		}
	}

	// 更新用户的信息
	isSuccess := updateUserInfo(c, user)
	if !isSuccess {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "Mysql修改用户信息错误！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil, "修改用户信息成功！"))
}

//
// UpdateUserById
//  @Description: 管理员修改指定id的用户, 需要请求参数: id
// 				 并且需要JSON参数: id(必须),user_account(必须),user_name,gender,phone,email,avatar_url
//  @param c
//
func UpdateUserById(c *gin.Context) {
	//验证id是否合法
	id := matchParamId(c)
	if id == "" {
		return
	}
	// 获取前端发送的用户信息
	var user models.UserDTO
	result := bindContextJson(c, &user)
	if !result {
		return
	}

	// 判断请求参数的id是否和获取到的用户id一致，不一致则直接返回
	if id != strconv.Itoa(int(user.ID)) {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "需要修改的用户参数和请求不合法！"))
		return
	}

	var userAccount string
	// 必须先获取用户原本的用户名，再进行判断用户是否已经更改了用户名
	affected := DB.Take(&models.User{}, id).Select("user_account").Scan(&userAccount).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusForbidden,
			utils.ResponseError(utils.ParamsError, "请求的id用户不存在！"))
		return
	}
	// 用户名发生了变化
	if userAccount != user.UserAccount {
		// 判断更改后的用户名是否存在
		exist := isUserAccountExist(c, user.UserAccount)
		if exist {
			return
		}
	}

	// 更新用户的信息
	isSuccess := updateUserInfo(c, user)
	if !isSuccess {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(utils.MysqlError, "Mysql修改用户信息错误！"))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil, "修改用户信息成功！"))
}

//
// ChangePasswordBySelf
//  @Description: 用户自己修改自己的密码,需要JSON参数: id,user_password,check_password (三个参数均必须)
//  @param c
//
func ChangePasswordBySelf(c *gin.Context) {
	// 获取需要修改的用户id、用户名、和修改后的密码
	var user models.UserChangePassword
	isSuccess := bindContextJson(c, &user)
	if !isSuccess {
		return
	}

	// 获取当前发起请求的用户是否和发送过来的json中的id是否一致
	userId := getContextValue(c, "userId")
	if userId == nil {
		return
	}
	if userId != strconv.Itoa(int(user.ID)) {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "身份验证失败！"))
		return
	}
	isSuccess = changePassword(c, user)
	if !isSuccess {
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil))
}

//
// ChangePasswordById
//  @Description: 管理员更改用户的密码,需要请求参数: id
//								需要JSON参数: id,user_password,check_password (三个参数均必须)
//  @param c
//
func ChangePasswordById(c *gin.Context) {
	// 获取请求的id
	id := matchParamId(c)
	if id == "" {
		return
	}
	// 获取需要修改的用户id、用户名、和修改后的密码
	var user models.UserChangePassword
	isSuccess := bindContextJson(c, &user)
	if !isSuccess {
		return
	}

	// 判断请求的id和发送的用户数据中的id是否一致
	if id != strconv.Itoa(int(user.ID)) {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "请求的id和发送的id不一致！"))
		return
	}

	isSuccess = changePassword(c, user)
	if !isSuccess {
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(nil))
}
