package factory

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenFactory struct {}
 
func NewAccessTokenFactory() *AccessTokenFactory {
	return &AccessTokenFactory{}
}

func (f *AccessTokenFactory) GetAccessToken(id string, email, role string, firstTimeLogin bool) (string, error) {
	accessSecret := os.Getenv("ACCESS_JWT_SECRET")

	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
		"email": email,
		"firstTimeLogin": firstTimeLogin,
		"role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

