package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

type TokenCache struct {
	RedisClient *redis.Client
}

func (tc *TokenCache) Get(key string) (*string, error) {
	value, err := tc.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		log.Println("Error in getting token from cache", err)
		return nil, err
	}
	return &value, nil
}

// Set a key with a value in redis token DB
func (tc *TokenCache) Set(key string, value interface{}, ttl_in_hour int64) error {
	err := tc.RedisClient.Set(context.Background(), key, value, time.Duration(ttl_in_hour)*time.Minute).Err()
	if err != nil {
		log.Println("Error is cache set for token", err)
		return err
	}
	return nil
}

// Invalidate a key in redis token DB
func (tc *TokenCache) InvalidateKey() {
	// TODO :: Update logic
}
