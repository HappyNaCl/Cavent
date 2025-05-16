package repo

import "github.com/HappyNaCl/Cavent/backend/internal/domain/model"

type TokenRepository interface {
	CheckToken(token string) (bool, error)
	IssueToken(token *model.RefreshToken) (*string, error)
	RevokeToken(token string) error
}

