package command

type UpdateUserInterestCommand struct {
	UserId      string   `json:"userId"`
	CategoryIds []string `json:"categoryIds"`
}

type UpdateUserInterestCommandResult struct {}