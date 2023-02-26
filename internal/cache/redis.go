package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

// NewRedisCache function used to connect with redis server and return new instance of `RedisCache`
func NewRedisCache(hostname, password, port string) (Cache, error) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", hostname, port),
		Password: password,
		DB:       0,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &RedisCache{
		client: rdb,
	}, nil
}

// Get Used to retrieve a cached value.
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key '%s' was not found on cache, maybe the same was not available", key)
	}
	if err != nil {
		return "", err
	}

	return val, nil
}

// Set Used to store values based on key on the cache context.
func (r *RedisCache) Set(ctx context.Context, key string, value string, ttl float64) error {
	exists, _ := r.client.Exists(ctx, key).Result()
	if exists == 1 {
		return fmt.Errorf("this key aready exits on cache")
	}

	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
