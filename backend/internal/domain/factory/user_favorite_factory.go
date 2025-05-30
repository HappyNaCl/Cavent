package factory

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type UserFavoriteFactory struct{}

func NewUserFavoriteFactory() *UserFavoriteFactory {
	return &UserFavoriteFactory{}
}

func (f *UserFavoriteFactory) GetUserFavorite(userId string, eventId uuid.UUID) *model.UserFavorite {
	return &model.UserFavorite{
		UserId:  userId,
		EventId: eventId,
	}
}