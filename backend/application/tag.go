package application

import (
	"github.com/HappyNaCl/Cavent/backend/config"
	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"github.com/HappyNaCl/Cavent/backend/infrastructure/persistence"
)

func GetAllTagsWithType() ([]*model.TagType, error) {
	tags, err := persistence.TagRepository(config.Database).GetAllTagsWithType()
	if err != nil {
		return nil, err
	}
	return tags, nil
}