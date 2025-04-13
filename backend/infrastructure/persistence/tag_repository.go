package persistence

import (
	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"github.com/HappyNaCl/Cavent/backend/domain/repository"
	"gorm.io/gorm"
)

type TagRepostioryImpl struct {
	Conn *gorm.DB
}

func TagRepository(conn *gorm.DB) repository.TagRepository{
	return &TagRepostioryImpl{Conn: conn}
}

func (repo *TagRepostioryImpl) GetAllTagsWithType() ([]*model.TagType, error) {
	var tagTypes []*model.TagType
	err := repo.Conn.
	Select("id", "name").
	Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "tag_type_id")
	}).
	Find(&tagTypes).Error
	if err != nil {
		return nil, err
	}
	return tagTypes, nil
}

func (repo *TagRepostioryImpl) GetTagTypes() ([]*model.TagType, error) {
	var tagTypes []*model.TagType
	err := repo.Conn.Find(&tagTypes).Error
	if err != nil {
		return nil, err
	}
	return tagTypes, nil
}

func (repo *TagRepostioryImpl) GetTags() ([]*model.Tag, error) {
	var tags []*model.Tag
	err := repo.Conn.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}