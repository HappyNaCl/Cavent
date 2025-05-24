package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func ConnectRedis() (*redis.Client, error) {
	url := os.Getenv("REDIS_URL")
	client := redis.NewClient(&redis.Options{
		Addr: url,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Sugar().Errorf("Error connecting to redis: %v", err)
		return nil, err
	}

	return client, nil
}