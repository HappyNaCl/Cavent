package main

import (
	"flag"

	"github.com/HappyNaCl/Cavent/backend/config"
	migration "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/postgresdb/migration"
	seeder "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/postgresdb/seeder"
	v1 "github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	_ "github.com/HappyNaCl/Cavent/backend/docs"
	swaggerFiles "github.com/swaggo/files"
)

// @title Cavent API
// @version 1.0
// @description This is the Cavent backend API documentation.
// @BasePath /api/v1
func main(){
	r := v1.Init()

	err := godotenv.Load()
    if err != nil {
        panic("Failed to load .env file")
    }

	config.SetupLogger()

	db, err := config.ConnectDb()
	if err != nil {
		zap.L().Sugar().Errorf("Error connecting to db: %v", err)
		panic(err)
	}

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}