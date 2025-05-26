package common

import (
	"github.com/google/uuid"
)

type CategoryResult struct {
	Id    uuid.UUID   `json:"id"`
	Name  string      `json:"name"`
}