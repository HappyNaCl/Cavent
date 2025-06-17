package request

import "github.com/HappyNaCl/Cavent/backend/internal/app/command"

type UpdateUserPasswordRequest struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"currentPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (r *UpdateUserPasswordRequest) ToCommand() *command.UpdateUserPasswordCommand {
	return &command.UpdateUserPasswordCommand{
		NewPassword: r.NewPassword,
		OldPassword: r.OldPassword,
	}
}