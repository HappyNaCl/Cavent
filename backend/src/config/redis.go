package config

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func ConnectRedis() error {
	url := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(url)
	if err != nil {
		return err
	}

	client := redis.NewClient(opt)

	RedisClient = client

	log.Printf("[INFO] Redis successfully connected!")

	return nil
}