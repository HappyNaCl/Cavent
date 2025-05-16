package factory

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

type RefreshTokenFactory struct {}

func generateToken() (string, error) {
	byteLength := 60

	b := make([]byte, byteLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (r *RefreshTokenFactory) GetToken(userId string) (*model.RefreshToken, error) {
	tokenStr, err := generateToken()
	if err != nil {
		return nil, err
	}

	return &model.RefreshToken{
		UserId: userId,
		Token: tokenStr,
	}, nil
}