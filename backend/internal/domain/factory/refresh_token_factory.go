package factory

import (
	"os"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenFactory struct {}

func NewRefreshTokenFactory() *RefreshTokenFactory {
	return &RefreshTokenFactory{}
}

func (f *RefreshTokenFactory) GetRefreshToken(user *model.User) (string, error)  {
	refreshSecret := os.Getenv("REFRESH_JWT_SECRET")

	claims := jwt.MapClaims{
		"Sub" : user.Id.String(),
		"Exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
		"Iat": time.Now().Unix(),
		"Email" : user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(refreshSecret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}



