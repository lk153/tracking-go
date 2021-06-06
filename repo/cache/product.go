package cache

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
	productCachePrefix = "prod_cache"
)

//CacheInteface ...
type CacheInteface interface {
	Get(ctx context.Context, key string) (*repo.ProductModel, error)
	Set(ctx context.Context, key string, value *repo.ProductModel) error
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
	cmd := pipeline.Get(ctx, fmt.Sprintf("%v_%v", productCachePrefix, key))
	_, err := pipeline.Exec(ctx)
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, fmt.Errorf("getCache %w", err)
	}

	buf, err := cmd.Bytes()
	if err != nil {
		return nil, fmt.Errorf("read bytes has error: %v", err)
	}

	if buf == nil {
		return nil, nil
	}

	var p *repo.ProductModel
	if err = json.Unmarshal(buf, &p); err != nil {
		return nil, fmt.Errorf("unmarshal product: %v", err)
	}

	return p, nil
}

//Set ...
func (r *RedisCache) Set(ctx context.Context, key string, value *repo.ProductModel) error {
	pipeline := r.rdb.Pipeline()
	cmds := []*redis.StatusCmd{}
	cacheData, err := json.Marshal(&value)
	if err != nil {
		return fmt.Errorf("set cache marshal error: %v", err)
	}

	cmds = append(cmds, pipeline.Set(ctx, fmt.Sprintf("%v_%v", productCachePrefix, key), string(cacheData), productFullInfoTTL))
	_, err = pipeline.Exec(ctx)
	if err != nil {
		return fmt.Errorf("set cache exec error: %v", err)
	}

	for _, cmd := range cmds {
		if cmd.Err() != nil {
			return fmt.Errorf("set product cache error: %v", cmd.Err())
		}
	}

	return nil
}
