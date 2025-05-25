package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type LoginUserCommand struct {
	Email    string
	Password string
}

type LoginUserCommandResult struct {
	Result *common.UserResult
}