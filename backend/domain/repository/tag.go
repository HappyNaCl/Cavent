package repository

import "github.com/HappyNaCl/Cavent/backend/domain/model"

type TagRepository interface{
	GetAllTagsWithType() ([]*model.TagType, error)
	GetTagTypes() ([]*model.TagType, error)
	GetTags() ([]*model.Tag, error)
}