package factory

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessTokenFactory struct {}
 
func NewAccessTokenFactory() *AccessTokenFactory {
	return &AccessTokenFactory{}
}

func (f *AccessTokenFactory) GetAccessToken(id uuid.UUID, email, role string, firstTimeLogin bool) (string, error) {
	accessSecret := os.Getenv("ACCESS_JWT_SECRET")

	claims := jwt.MapClaims{
		"Sub": id.String(),
		"Exp": time.Now().Add(10 * time.Minute).Unix(),
		"Iat": time.Now().Unix(),
		"Email": email,
		"FirstTimeLogin": firstTimeLogin,
		"Role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

