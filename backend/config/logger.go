package config

import (
	"os"

	"go.uber.org/zap"
)

func SetupLogger() *zap.Logger {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	if appEnv == "development" {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		zap.ReplaceGlobals(logger)
		return nil
	}
	
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)
	return logger
}