package factory

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RefreshTokenFactory struct {}

func NewRefreshTokenFactory() *RefreshTokenFactory {
	return &RefreshTokenFactory{}
}

func (f *RefreshTokenFactory) GetRefreshToken(id uuid.UUID, email string) (string, error)  {
	refreshSecret := os.Getenv("REFRESH_JWT_SECRET")

	claims := jwt.MapClaims{
		"Sub" : id.String(),
		"Exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
		"Iat": time.Now().Unix(),
		"Email" : email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(refreshSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}



