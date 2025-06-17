package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetUserProfileCommand struct {
	UserId string `json:"userId"`
}

type GetUserProfileCommandResult struct {
	Result common.UserProfileResult `json:"result"`
}