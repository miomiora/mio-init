package api

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user-center/config"
	"user-center/models"
)

var DB *gorm.DB
var Conn redis.Conn

// 连接数据库
func init() {
	connectMysql()
	connectRedis()
}

func connectMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Address,
		config.Config.Mysql.Dbname,
		config.Config.Mysql.Timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("连接Mysql数据库失败, error=" + err.Error())
		return
	}
	// 连接成功
	fmt.Println("[Success] Mysql数据库连接成功！！！")
	DB = conn
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("创建表失败！")
	}
}

func connectRedis() {
	c, err := redis.Dial("tcp", config.Config.Redis.Address)
	if err != nil {
		fmt.Println("连接Redis数据库失败！" + err.Error())
	}
	fmt.Println("[Success] Redis数据库连接成功！！！")
	Conn = c
}
