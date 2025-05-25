package request

import "github.com/HappyNaCl/Cavent/backend/internal/app/command"

type LoginRequest struct {
	Email    string   `json:"email" form:"email" binding:"required,email"`
	Password string   `json:"password" form:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe" form:"rememberMe"`
}

func (r LoginRequest) ToCommand() *command.LoginUserCommand {
	return &command.LoginUserCommand{
		Email:    r.Email,
		Password: r.Password,
	}
}