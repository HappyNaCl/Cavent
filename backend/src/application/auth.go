package application

import (
	"errors"
	"os"
	"time"

	"github.com/HappyNaCl/Cavent/src/config"
	"github.com/HappyNaCl/Cavent/src/domain"
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

func RegisterOrLoginOauthUser(user goth.User, provider string) (*domain.User, error){
	return persistence.UserRepository(config.Database).RegisterOrLoginOauthUser(user, provider)
}

func LoginUser(email string, password string) (*domain.User, error){
	user, err := persistence.UserRepository(config.Database).FindByProviderID(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}