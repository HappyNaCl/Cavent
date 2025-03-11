package application

import (
	"os"
	"time"

	"github.com/HappyNaCl/Cavent/src/config"
	"github.com/HappyNaCl/Cavent/src/infrastructure/persistence"
	"github.com/golang-jwt/jwt"
	"github.com/markbates/goth"
)

func GenerateJWT(email string, provider string) (string, error){
	claims := jwt.MapClaims{
		"email": email,
		"provider": provider,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	byteSecret := []byte(jwtSecret)
	return token.SignedString(byteSecret)
}

func RegisterOrLoginUser(user goth.User, provider string) (bool, error){
	return persistence.UserRepository(config.Database).RegisterOrLoginUser(user, provider)
}