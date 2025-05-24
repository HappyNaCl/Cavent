package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type RegisterUserCommand struct {
	Name 			string
	Email 			string
	Password 		string
	ConfirmPassword string
	InviteCode      *string
}

type RegisterUserCommandResult struct {
	Result *common.UserResult
}