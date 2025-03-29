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

var (
	client *redis.Client
	ctx    context.Context
)

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

	ctx = context.Background()
	_, err = client.Ping(ctx).Result()
	zap.L().Info("[dao redis Init] ping redis client failed ", zap.Error(err))
	return
}

func (redisCore) Close() {
	err := client.Close()
	zap.L().Info("[dao redis Close] close the redis connect failed ", zap.Error(err))
}
