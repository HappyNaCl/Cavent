package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FavoriteGormRepository struct {
	db *gorm.DB
}

func (i *FavoriteGormRepository) IsFavorited(userId string, eventIds []uuid.UUID) (map[uuid.UUID]bool, error) {
	var favorites []model.UserFavorite
	err := i.db.
		Where("user_id = ? AND event_id IN ?", userId, eventIds).
		Find(&favorites).Error
	if err != nil {
		return nil, err
	}

	result := make(map[uuid.UUID]bool, len(eventIds))

	for _, id := range eventIds {
		result[id] = false
	}

	for _, fav := range favorites {
		result[fav.EventId] = true
	}

	return result, nil
}

func NewFavoriteGormRepo(db *gorm.DB) *FavoriteGormRepository {
	return &FavoriteGormRepository{
		db: db,
	}
}

func (i *FavoriteGormRepository) FavoriteEvent(fav *model.UserFavorite) (int64, error) {
	err := i.db.Clauses(clause.OnConflict{DoNothing: true}).Create(fav).Error
	if err != nil {
		return 0, err
	}

	var count int64
	err = i.db.Model(&model.UserFavorite{}).
		Where("event_id = ?", fav.EventId).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	err = i.db.Model(&model.Event{}).
		Where("id = ?", fav.EventId).
		Update("favorite_count", count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// UninterestEvent implements repo.InterestRepository.
func (i *FavoriteGormRepository) UnfavoriteEvent(fav *model.UserFavorite) (int64, error) {
	err := i.db.Where("user_id = ? AND event_id = ?", fav.UserId, fav.EventId).
		Delete(&model.UserFavorite{}).Error
	if err != nil {
		return 0, err
	}

	var count int64
	err = i.db.Model(&model.UserFavorite{}).
		Where("event_id = ?", fav.EventId).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	err = i.db.Model(&model.Event{}).
		Where("id = ?", fav.EventId).
		Update("favorite_count", count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func NewFavoriteRepository(db *gorm.DB) repo.FavoriteRepository {
	return &FavoriteGormRepository{
		db: db,
	}
}
