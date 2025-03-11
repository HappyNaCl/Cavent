package main

import (
	"context"

	"github.com/HappyNaCl/Cavent/src/config"
	"github.com/HappyNaCl/Cavent/src/interfaces"
	"github.com/joho/godotenv"
	"google.golang.org/appengine/log"
)

func main(){
	
	err := godotenv.Load("../.env")
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

	interfaces.Run(8080)
}