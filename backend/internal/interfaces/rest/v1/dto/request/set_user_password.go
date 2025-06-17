package request

import "github.com/HappyNaCl/Cavent/backend/internal/app/command"

type SetUserPasswordRequest struct {
	NewPassword string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (r *SetUserPasswordRequest) ToCommand() *command.SetPasswordCommand {
	return &command.SetPasswordCommand{
		NewPassword:     r.NewPassword,
	}
}