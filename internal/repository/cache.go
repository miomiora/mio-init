package repository

import (
	"context"
	"mio-init/internal/core"
	"time"
)

type cacheRepo struct {
}

var Cache = new(cacheRepo)

func (cacheRepo) Set(ctx context.Context, key string, value string, expire time.Duration) (result string, err error) {
	return core.Redis.GetClient().Set(ctx, key, value, expire).Result()
}

func (cacheRepo) SetKeepTll(ctx context.Context, key string, value string) (result string, err error) {
	return core.Redis.GetClient().Set(ctx, key, value, 0).Result()
}

func (cacheRepo) Get(ctx context.Context, key string) (result string, err error) {
	return core.Redis.GetClient().Get(ctx, key).Result()
}

func (cacheRepo) Del(ctx context.Context, key string) (result int64, err error) {
	return core.Redis.GetClient().Del(ctx, key).Result()
}
