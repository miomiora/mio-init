package repository

import (
	"context"
	"mio-init/internal/core"
	"mio-init/util"
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

func (cacheRepo) Exists(ctx context.Context, key string) (result int64, err error) {
	return core.Redis.GetClient().Exists(ctx, key).Result()
}

func (cacheRepo) Logout(ctx context.Context, refreshToken, accessToken string) error {

	pipeline := core.Redis.GetClient().Pipeline()
	pipeline.Del(ctx, util.GenRefreshKey(refreshToken))
	pipeline.Set(ctx, util.GenBlackListKey(accessToken), "1", util.GetRemainingTTL(accessToken))

	_, err := pipeline.Exec(ctx)

	return err
}

func (cacheRepo) RefreshToken(ctx context.Context, refreshToken, newToken string, userId int64) error {

	pipeline := core.Redis.GetClient().Pipeline()
	pipeline.Del(ctx, util.GenRefreshKey(refreshToken))
	pipeline.Set(ctx, util.GenRefreshKey(newToken), userId, util.RefreshTokenExpire)

	_, err := pipeline.Exec(ctx)

	return err
}
