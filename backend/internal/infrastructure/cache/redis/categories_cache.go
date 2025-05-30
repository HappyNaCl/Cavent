package rediscache

import (
	"context"
	"encoding/json"

	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/cache"
	"github.com/redis/go-redis/v9"
)

type CategoriesCache struct {
	redisClient *redis.Client
}

func NewCategoriesCache(redisClient *redis.Client) cache.CategoriesCache {
	return &CategoriesCache{
		redisClient: redisClient,
	}
}

const CACHE_CATEGORIES_KEY = "categories:all" 

func (c *CategoriesCache) GetAllCategories(ctx context.Context) ([]*common.CategoryResult, error) {
	data, err := c.redisClient.Get(ctx, CACHE_CATEGORIES_KEY).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err 
	}

	var categories []*common.CategoryResult
	err = json.Unmarshal([]byte(data), &categories)
	if err != nil {
		return nil, err 
	}

	return categories, nil
}

func (c* CategoriesCache) SetAllCategories(ctx context.Context, categories []*common.CategoryResult) error {
	if len(categories) == 0 {
		return nil 
	}

	data, err := json.Marshal(categories)
	if err != nil {
		return err 
	}

	err = c.redisClient.Set(ctx, CACHE_CATEGORIES_KEY, data, 0).Err()
	if err != nil {
		return err 
	}

	return nil
}

func (c *CategoriesCache) Invalidate(ctx context.Context) error {
	return c.redisClient.Del(ctx, CACHE_CATEGORIES_KEY).Err()
}