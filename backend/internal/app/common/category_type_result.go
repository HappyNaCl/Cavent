package common

type CategoryTypeResult struct {
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	Categories    []*CategoryResult  `json:"categories"`
}