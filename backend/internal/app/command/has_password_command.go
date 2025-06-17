package command

type HasPasswordCommand struct {
	UserId string `json:"userId"`
}

type HasPasswordCommandResult struct {
	HasPassword bool `json:"hasPassword"`
}