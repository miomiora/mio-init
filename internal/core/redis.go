package core

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"mio-init/config"
	"time"
)

type redisCore struct {
}

var Redis = new(redisCore)

var client *redis.Client

const (
	TokenPrefix  = "login:token:"
	TokenTimeout = time.Hour * 24
)

func (redisCore) Init(cfg *config.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // 密码
		DB:       cfg.Db,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err = client.Ping(context.Background()).Result()
	zap.L().Info("[repository redis Init] ping redis client failed ", zap.Error(err))
	return
}

func (redisCore) Close() {
	err := client.Close()
	zap.L().Info("[repository redis Close] close the redis connect failed ", zap.Error(err))
}

func (redisCore) GetClient() *redis.Client {
	return client
}
