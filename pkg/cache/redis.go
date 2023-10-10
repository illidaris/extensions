package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/illidaris/logger"
)

var (
	redisClient *redis.Client
)

func InitRedis(ctx context.Context) error {
	opt := DefaultRedisOption()
	e := redis.NewClient(&opt)
	res, err := e.Ping(ctx).Result()
	logger.InfoCtx(ctx, fmt.Sprintf("init redis %s,ping=>%s", opt.Addr, res))
	if err != nil {
		return err
	}
	redisClient = e
	return nil
}

func DefaultRedisOption() redis.Options {
	return redis.Options{
		// Addr:        fmt.Sprintf("%s:%d", config.GetString("redis.host"), config.GetInt("redis.port")),
		// DB:          config.GetInt("redis.db"),
		// Username:    config.GetString("redis.username"),
		// Password:    config.GetString("redis.password"),
		// DialTimeout: time.Duration(config.GetInt32("redis.timeout.dial")) * time.Millisecond,
	}
}

type RedisCache struct {
}

func NewRedisCache() *RedisCache {
	return &RedisCache{}
}

func (c *RedisCache) Get(key string) interface{} {
	v, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil
	}
	return v
}

func (c *RedisCache) TTL(key string) time.Duration {
	v, err := redisClient.TTL(context.Background(), key).Result()
	if err != nil {
		return time.Duration(0)
	}
	return v
}

func (c *RedisCache) Set(key string, val interface{}, timeout time.Duration) error {
	_, err := redisClient.Set(context.Background(), key, val, timeout).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisCache) SetNX(key string, val interface{}, timeout time.Duration) (bool, error) {
	return redisClient.SetNX(context.Background(), key, val, timeout).Result()
}

func (c *RedisCache) IsExist(key string) bool {
	v, err := redisClient.Exists(context.Background(), key).Result()
	if err != nil {
		return false
	}
	return v > 0
}

func (c *RedisCache) Delete(key string) error {
	_, err := redisClient.Del(context.Background(), key).Result()
	return err
}
