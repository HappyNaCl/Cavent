package command

type UpdateUserPasswordCommand struct {
	UserId      string
	OldPassword string
	NewPassword string
}

type UpdateUserPasswordCommandResult struct {}