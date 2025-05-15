package main

import (
	"context"
	"flag"

	"github.com/HappyNaCl/Cavent/backend/config"
	_ "github.com/HappyNaCl/Cavent/backend/docs"
	"github.com/HappyNaCl/Cavent/backend/interfaces"
	"github.com/joho/godotenv"
	"google.golang.org/appengine/log"
)

// @title Cavent API
// @version 1.0
// @description This is the Cavent backend API documentation.
// @BasePath /api/v1
func main(){
	
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	config.SetupOAuth()

	err = config.ConnectDatabase()
	if err != nil {
		log.Infof(context.Background(), "Error connecting to database: %v", err.Error())
		panic(err)
	}

	err = config.ConnectRedis()
	if err != nil {
		log.Infof(context.Background(), "Error connecting to redis: %v", err.Error())
		panic(err)
	}

	migrate := flag.Bool("migrate", false, "Run migrations")
	seed := flag.Bool("seed", false, "Run seeders")
	fresh := flag.Bool("fresh", false, "Run fresh migrations and seed")

	flag.Parse()

	if migrate != nil && *migrate {
		err = config.Migrate(config.Database)
		if err != nil {
			log.Infof(context.Background(), "Error migrating database: %v", err.Error())
			panic(err)
		}
	}

	if seed != nil && *seed {
		err = config.Seed(config.Database)
		if err != nil {
			log.Infof(context.Background(), "Error seeding database: %v", err.Error())
			panic(err)
		}
	}

	if fresh != nil && *fresh {
		err = config.Migrate(config.Database)
		if err != nil {
			log.Infof(context.Background(), "Error migrating database: %v", err.Error())
			panic(err)
		}
		err = config.Seed(config.Database)
		if err != nil {
			log.Infof(context.Background(), "Error seeding database: %v", err.Error())
			panic(err)
		}
	}

	interfaces.Run(8080)
}