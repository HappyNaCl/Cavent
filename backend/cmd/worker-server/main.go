package main

import (
	"log"
	"os"

	"github.com/HappyNaCl/Cavent/backend/config"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/handler"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	logger := config.SetupLogger()
	defer logger.Sync()

	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	redisAddr := os.Getenv("REDIS_URL")
	
	eventViewHandler := handler.NewEventViewedHandler(db)
	queue.StartWorker(redisAddr, eventViewHandler)
}
