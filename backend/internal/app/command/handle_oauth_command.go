package command

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/markbates/goth"
)

type HandleOAuthCommand struct {
	Provider 	string 
	User 		goth.User 
}

type HandleOAuthCommandResult struct {
	Result *common.UserResult
}