package utils

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	RedisClient *redis.Client
}

// NewCacheService creates a new instance of CacheService
func NewCacheService(redisClient *redis.Client) *CacheService {
	return &CacheService{RedisClient: redisClient}
}

// GetCache fetches the value for a given key from Redis
func (c *CacheService) GetCache(ctx context.Context, key string) (bool, interface{}, error) {
	cachedData, err := c.RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// Key does not exist
		return false, nil, nil
	} else if err != nil {
		// Other Redis errors
		return false, nil, err
	}

	var data interface{}

	err = json.Unmarshal([]byte(cachedData), &data)
	if err != nil {
		return false, nil, err
	}

	return true, data, nil
}

// SetCache stores a value in Redis for a given key with an expiration time
func (c *CacheService) SetCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	marshaledData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.RedisClient.Set(ctx, key, marshaledData, expiration).Err()
}
