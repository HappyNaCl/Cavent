package rediscache

import (
	"context"
	"fmt"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/cache"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type EventFavoriteCache struct {
	redisClient *redis.Client
}

// DecrementEventFavoriteCount implements cache.EventFavoriteCache.
func (e *EventFavoriteCache) DecrementEventFavoriteCount(ctx context.Context, eventID uuid.UUID) (int64, error) {
	key := fmt.Sprintf(CACHE_KEY, eventID)

	count, err := e.redisClient.Decr(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	if count < 0 {
		if err := e.redisClient.Set(ctx, key, 0, 0).Err(); err != nil {
			return 0, err
		}
		return 0, nil
	}

	return count, nil
}

// GetEventFavoriteCount implements cache.EventFavoriteCache.
func (e *EventFavoriteCache) GetEventFavoriteCount(ctx context.Context, eventID uuid.UUID) (int64, error) {
	key := fmt.Sprintf(CACHE_KEY, eventID)

	count, err := e.redisClient.Get(ctx, key).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}

// IncrementEventFavoriteCount implements cache.EventFavoriteCache.
func (e *EventFavoriteCache) IncrementEventFavoriteCount(ctx context.Context, eventID uuid.UUID) (int64, error) {
	key := fmt.Sprintf(CACHE_KEY, eventID)

	count, err := e.redisClient.Incr(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}
	
	return count, nil
}

const CACHE_KEY = "event:%s:favorite_count"

func NewEventFavoriteCache(redisClient *redis.Client) cache.EventFavoriteCache {
	return &EventFavoriteCache{
		redisClient: redisClient,
	}
}
