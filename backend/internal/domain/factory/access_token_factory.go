package factory

import (
	"os"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenFactory struct {}
 
func NewAccessTokenFactory() *AccessTokenFactory {
	return &AccessTokenFactory{}
}

func (f *AccessTokenFactory) GetAccessToken(user *model.User) (string, error) {
	accessSecret := os.Getenv("ACCESS_JWT_SECRET")

	claims := jwt.MapClaims{
		"Sub": user.Id.String(),
		"Exp": time.Now().Add(10 * time.Minute).Unix(),
		"Iat": time.Now().Unix(),
		"Email": user.Email,
		"FirstTimeLogin": user.FirstTimeLogin,
		"Role": user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(accessSecret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

