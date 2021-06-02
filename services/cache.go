package services

import (
	"context"
	"fmt"
	"time"

	"factory/exam/repo"

	"encoding/json"

	"github.com/go-redis/redis/v8"
)

const (
	productFullInfoTTL = 3 * 24 * time.Hour
)

//CacheInteface ...
type CacheInteface interface {
	Get(ctx context.Context, key string) (*repo.ProductModel, error)
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
func (r *RedisCache) Get(ctx context.Context, key string) (*repo.ProductModel, error) {
	pipeline := r.rdb.Pipeline()
	cmds := []*redis.StringCmd{}
	cmds = append(cmds, pipeline.Get(ctx, key))
	for _, raw := range cmds {
		buf, err := raw.Bytes()
		if err != nil || buf == nil {
			if err != redis.Nil {
				fmt.Errorf("Read bytes has error %v", err)
			}
			continue
		}

		var p repo.ProductModel
		if err = json.Unmarshal(buf, &p); err != nil {
			fmt.Errorf("Unmarshal Product", err)
			continue
		}

	}

	return nil, nil
}

//Set ...
func (r *RedisCache) Set(ctx context.Context, key string, value string) error {
	pipeline := r.rdb.Pipeline()
	cmds := []*redis.StatusCmd{}
	cmds = append(cmds, pipeline.Set(ctx, key, string(value), productFullInfoTTL))

	_, err := pipeline.Exec(ctx)

	for _, cmd := range cmds {
		if cmd.Err() != nil {
			fmt.Errorf("Set Product Cache Error", cmd.Err())
		}
	}

	return err
}
