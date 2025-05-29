package rediscache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/cache"
	"github.com/redis/go-redis/v9"
)

type UserInterestCache struct {
	redis *redis.Client
}

const CACHE_USER_INTEREST_KEY = "user:%s:interest"

func NewUserInterestCache(redis *redis.Client) cache.UserInterestCache {
	return &UserInterestCache{
		redis: redis,
	}
}

// GetUserInterest fetches cached CategoryResults
func (c *UserInterestCache) GetUserInterest(ctx context.Context, userId string) ([]*common.CategoryResult, error) {
	key := fmt.Sprintf(CACHE_USER_INTEREST_KEY, userId)

	data, err := c.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user interest: %w", err)
	}

	var interests []*common.CategoryResult
	if err := json.Unmarshal([]byte(data), &interests); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user interest: %w", err)
	}

	return interests, nil
}

// SetUserInterest caches CategoryResults
func (c *UserInterestCache) SetUserInterest(ctx context.Context, userId string, categories []*common.CategoryResult) error {
	key := fmt.Sprintf(CACHE_USER_INTEREST_KEY, userId)

	if len(categories) == 0 {
		return nil 
	}

	data, err := json.Marshal(categories)
	if err != nil {
		return fmt.Errorf("failed to marshal user interest: %w", err)
	}

	err = c.redis.Set(ctx, key, data, 10 * time.Minute).Err()
	if err != nil {
		return fmt.Errorf("failed to set user interest: %w", err)
	}

	return nil
}

func (c *UserInterestCache) Invalidate(ctx context.Context, userId string) error {
	key := fmt.Sprintf(CACHE_USER_INTEREST_KEY, userId)

	err := c.redis.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to invalidate user interest cache: %w", err)
	}

	return nil
}