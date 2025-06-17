package main

import (
	"os"

	"github.com/HappyNaCl/Cavent/backend/config"

	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue"
	v1 "github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	_ "github.com/HappyNaCl/Cavent/backend/docs"
)

// @title Cavent API
// @version 1.0
// @description This is the Cavent backend API documentation.
// @BasePath /api/v1
func main(){
	err := godotenv.Load()
    if err != nil {
        panic("Failed to load .env file")
    }

	logger := config.SetupLogger()
	defer logger.Sync()

	config.SetupOAuth()

	db, err := config.ConnectDb()
	if err != nil {
		zap.L().Sugar().Errorf("Error connecting to db: %v", err)
		panic(err)
	}

	redis, err := config.ConnectRedis()
	if err != nil {
		zap.L().Sugar().Errorf("Error connecting to redis: %v", err)
		panic(err)
	}

	redisAddr := os.Getenv("REDIS_URL")

	// Setup Asynq client
	client, err := queue.InitClient(redisAddr)
	if err != nil {
		panic(err)
	}

	// Setup Gin router
	r := v1.Init(db, redis, client)
	r.Run(":8080")

	defer redis.Close()
}