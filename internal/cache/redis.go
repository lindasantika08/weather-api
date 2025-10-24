package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{Client: rdb}
}

func (r *RedisCache) Set(key string, value string, expiration time.Duration) error {
	return r.Client.Set(Ctx, key, value, expiration).Err()
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.Client.Get(Ctx, key).Result()
}
