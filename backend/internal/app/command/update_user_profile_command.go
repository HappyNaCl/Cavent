package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type UpdateUserProfileCommand struct {
	UserId      string  `json:"userId"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	PhoneNumber *string `json:"phoneNumber"`
	Address     *string `json:"address"`
	AvatarBytes []byte  `json:"avatarBytes,omitempty"`
	AvatarExt   *string `json:"avatarExt,omitempty"`
}

type UpdateUserProfileCommandResult struct {
	Result      common.UserProfileResult `json:"result"`
}