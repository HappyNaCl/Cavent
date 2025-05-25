package mapper

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

func NewUserResultFromRegisteredUser(user *model.User) *common.UserResult {
	return NewUserResultFromUser(user)
}

func NewUserResultFromLoginUser(user *model.User) *common.UserResult {
	return NewUserResultFromUser(user)
}

func NewUserResultFromUser(user *model.User) *common.UserResult {
	if user == nil {
		return nil
	}

	return &common.UserResult{
		Id:             user.Id,
		CampusId:       user.CampusId,
		Provider:       user.Provider,
		ProviderId:     user.ProviderId,
		Email:          user.Email,
		Name:           user.Name,
		AvatarUrl:      user.AvatarUrl,
		Description:    user.Description,
		Role:           user.Role,
	}
}