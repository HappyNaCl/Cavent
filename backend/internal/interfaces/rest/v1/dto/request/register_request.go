package request

import "github.com/HappyNaCl/Cavent/backend/internal/app/command"

type RegisterRequest struct {
	Name 			string `json:"name" form:"name" binding:"required"`
	Password 		string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" binding:"required"`
	Email 			string `json:"email" form:"email" binding:"required"`
	InviteCode 		*string `json:"inviteCode" form:"inviteCode" binding:"required"`
}

func (r RegisterRequest) ToCommand() *command.RegisterUserCommand {
	return &command.RegisterUserCommand{
		Name: r.Name,
		Email: r.Email,
		Password: r.Password,
		ConfirmPassword: r.ConfirmPassword,
		InviteCode: r.InviteCode,
	}
}