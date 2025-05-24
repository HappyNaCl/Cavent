package request

import "github.com/HappyNaCl/Cavent/backend/internal/app/command"

type RegisterRequest struct {
	Name 			string `json:"name" form:"name"`
	Password 		string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword"`
	Email 			string `json:"email" form:"email"`
	InviteCode 		*string `json:"inviteCode" form:"inviteCode"`
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