package core

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mio-init/config"
	"mio-init/internal/model"
)

var db *gorm.DB

type mysqlCore struct {
}

var MySQL = new(mysqlCore)

func (mysqlCore) Init(cfg *config.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		zap.L().Error("[dao mysql Init] connect mysql error ", zap.Error(err))
		return
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		zap.L().Info("[dao mysql Init] create table failed ", zap.Error(err))
		return err
	}

	conn, err := db.DB()
	if err != nil {
		zap.L().Info("[dao mysql Init] get sql instance failed ", zap.Error(err))
		return err
	}
	conn.SetMaxOpenConns(cfg.MaxOpenConn)
	conn.SetMaxIdleConns(cfg.MaxIdleConn)
	return
}

func (mysqlCore) Close() {
	conn, err := db.DB()
	zap.L().Info("[dao mysql Close] get sql instance failed ", zap.Error(err))
	err = conn.Close()
	zap.L().Info("[dao mysql Close] close the mysql connect failed ", zap.Error(err))
}

func (mysqlCore) GetDB() *gorm.DB {
	return db
}
