package services

import (
	"context"

	"github.com/go-redis/redis/v8"
)

//CacheInteface ...
type CacheInteface interface {
	Get(ctx context.Context, key string) error
	Set(ctx context.Context, key string, value string) error
}

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

var _ CacheInteface = &RedisCache{}

type RedisCache struct {
	rdb *redis.Client
}

//NewRedisCacheRepo ...
func NewRedisCacheRepo() *RedisCache {
	return &RedisCache{
		rdb: redisClient,
	}
}

//Get ...
func (r *RedisCache) Get(ctx context.Context, key string) error {
	return nil
}

//Set ...
func (r *RedisCache) Set(ctx context.Context, key string, value string) error {
	return nil
}
