package command

type SetPasswordCommand struct {
	UserId string
	NewPassword string
}

type SetPasswordCommandResult struct {
}