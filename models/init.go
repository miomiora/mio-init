package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user-center/config"
)

var DB *gorm.DB

// 连接数据库
func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.Dbname,
		config.Config.Mysql.Timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("连接数据库失败, error=" + err.Error())
		return
	}
	// 连接成功
	fmt.Println("MySQL数据库连接成功！！！")
	DB = db
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("创建表失败！")
	}
}
