package main

import (
	"context"
	"errors"
	"fmt"
	"mio-init/config"
	"mio-init/dao/mysql"
	"mio-init/dao/redis"
	"mio-init/logger"
	"mio-init/router"
	"mio-init/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// @title mio-init
// @version 1.0
// @description Go Web 开发脚手架

// @contact.name miomiora
// @contact.url https://github.com/miomiora

// @host 127.0.0.1:8081
// @BasePath /api/v1
func main() {
	// 1、加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("init config error : %s \n", err)
		return
	}
	// 2、初始化日志
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("init logger error : %s \n", err)
		return
	}
	defer zap.L().Sync()
	// 3、初始化MySQL
	if err := mysql.Init(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql error : %s \n", err)
		return
	}
	defer mysql.Close()
	// 4、初始化Redis
	if err := redis.Init(config.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis error  %s \n", err)
		return
	}
	defer redis.Close()
	if err := util.Init(config.Conf.StartTime, config.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake error  %s \n", err)
		return
	}
	// 5、注册路由
	r := router.Setup()
	// 6、启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Info("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)                      // 创建一个接收信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
