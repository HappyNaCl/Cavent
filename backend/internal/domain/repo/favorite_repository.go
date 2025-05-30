package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type FavoriteRepository interface {
	FavoriteEvent(fav *model.UserFavorite) (int64, error)
	UnfavoriteEvent(fav *model.UserFavorite) (int64, error)
	IsFavorited(userId string, eventId []uuid.UUID) (map[uuid.UUID]bool, error)
}