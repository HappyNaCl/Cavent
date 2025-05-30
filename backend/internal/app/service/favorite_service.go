package service

import (
	"context"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/cache"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	rediscache "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/cache/redis"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type FavoriteService struct {
	eventFavoriteCache cache.EventFavoriteCache
	favoriteRepo repo.FavoriteRepository
}

func NewFavoriteService(db *gorm.DB, redis *redis.Client) *FavoriteService {
	return &FavoriteService{
		eventFavoriteCache: rediscache.NewEventFavoriteCache(redis),
		favoriteRepo: postgresdb.NewFavoriteRepository(db),	
	}
}

func (fs *FavoriteService) FavoriteEvent(ctx context.Context, com *command.FavoriteEventCommand) (*command.FavoriteEventCommandResult, error) {
	factory := factory.NewUserFavoriteFactory()
	fav := factory.GetUserFavorite(com.UserId, com.EventId)


	count, err := fs.favoriteRepo.FavoriteEvent(fav)
	if err != nil {
		return nil, err
	}
	
	_, err = fs.eventFavoriteCache.IncrementEventFavoriteCount(ctx, com.EventId)
	if err != nil {
		return nil, err
	}

	return &command.FavoriteEventCommandResult{
		Result: count,
	}, nil
}

func (fs *FavoriteService) UnfavoriteEvent(ctx context.Context, com *command.UnfavoriteEventCommand) (*command.UnfavoriteEventCommandResult, error) {
	factory := factory.NewUserFavoriteFactory()
	fav := factory.GetUserFavorite(com.UserId, com.EventId)

	count, err := fs.favoriteRepo.UnfavoriteEvent(fav)
	if err != nil {
		return nil, err
	}

	_, err = fs.eventFavoriteCache.DecrementEventFavoriteCount(ctx, com.EventId)
	if err != nil {
		return nil, err
	}

	return &command.UnfavoriteEventCommandResult{
		Result: count,
	}, nil
}