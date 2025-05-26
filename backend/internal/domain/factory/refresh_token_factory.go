package factory

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenFactory struct {}

func NewRefreshTokenFactory() *RefreshTokenFactory {
	return &RefreshTokenFactory{}
}

func (f *RefreshTokenFactory) GetRefreshToken(id string) (string, error)  {
	refreshSecret := os.Getenv("REFRESH_JWT_SECRET")

	claims := jwt.MapClaims{
		"sub" : id,
		"exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(refreshSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}



