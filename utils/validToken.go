package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"user-center/models"
)

//
// ValidToken
//  @Description: 验证token是否有效，有效的话继续验证用户权限
//  @param c
//  @param conn
//  @param db
//  @return int		返回-1表示验证token失败，0表示普通用户，1表示管理员
//  @return error
//
func ValidToken(c *gin.Context, conn redis.Conn, db *gorm.DB) (int, error) {
	// 从请求头中获取Token, 没有token就直接返回
	token := c.GetHeader("Authorization")
	if token == "" {
		return ROLE_UNDEFINED, nil
	}
	tokenKey := TOKEN_PREIX + token
	// 从redis中进行查询
	result, err := redis.Strings(conn.Do("HMGET", tokenKey, "client_ip", "id"))
	// 没查到就会报err
	if err != nil {
		return ROLE_UNDEFINED, err
	}
	// 判断ip是否和当前客户端请求的ip一致
	if result[0] != c.ClientIP() {
		return ROLE_UNDEFINED, nil
	}
	// 判断是否获取id
	if result[1] == "" {
		return ROLE_UNDEFINED, nil
	}
	// 获取到了id则根据id进行查询用户的role
	var role int
	affected := db.Select("role").
		Take(&models.User{}, "id = ?", result[1]).Scan(&role).RowsAffected
	if affected == 0 {
		c.JSON(500, gin.H{"message": "没查到token中id的用户"})
		return ROLE_UNDEFINED, nil
	}
	// 判断用户权限
	if role == ROLE_USER {
		return ROLE_USER, nil
	}
	if role == ROLE_ADMIN {
		return ROLE_ADMIN, nil
	}
	return ROLE_UNDEFINED, nil
}
