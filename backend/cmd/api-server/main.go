package main

import (
	"flag"
	"os"

	"github.com/HappyNaCl/Cavent/backend/config"
	migration "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb/migration"
	seeder "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb/seeder"
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
    
	migrate := flag.Bool("migrate", false, "Run migrations")
	seed := flag.Bool("seed", false, "Run seeders")
	fresh := flag.Bool("fresh", false, "Run fresh migrations and seed")

	flag.Parse()

	if migrate != nil && *migrate {
		err := migration.Migrate(db)
		if err != nil {
			zap.L().Sugar().Infof("Error migrating database: %v", err.Error())
			panic(err)
		}
	}

	if seed != nil && *seed {
		err := seeder.Seed(db)
		if err != nil {
			zap.L().Sugar().Infof("Error seeding database: %v", err.Error())
			panic(err)
		}
	}

	if fresh != nil && *fresh {
		err := migration.Migrate(db)
		if err != nil {
			zap.L().Sugar().Infof("Error migrating database: %v", err.Error())
			panic(err)
		}
		err = seeder.Seed(db)
		if err != nil {
			zap.L().Sugar().Infof("Error seeding database: %v", err.Error())
			panic(err)
		}
	}


	r.Run(":8080")
}