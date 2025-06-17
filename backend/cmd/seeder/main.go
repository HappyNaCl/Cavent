package main

import (
	"context"
	"flag"

	"github.com/HappyNaCl/Cavent/backend/config"
	migration "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb/migration"
	seeder "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb/seeder"
	"go.uber.org/zap"
)

func main() {
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
		
		err = redis.FlushAll(context.Background()).Err()
		if err != nil {
			zap.L().Sugar().Infof("Error flushing redis: %v", err.Error())
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
		err = redis.FlushAll(context.Background()).Err()
		if err != nil {
			zap.L().Sugar().Infof("Error flushing redis: %v", err.Error())
			panic(err)
		}
		err = seeder.Seed(db)
		if err != nil {
			zap.L().Sugar().Infof("Error seeding database: %v", err.Error())
			panic(err)
		}
	}


	defer redis.Close()
}