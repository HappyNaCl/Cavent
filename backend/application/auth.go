package application

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/HappyNaCl/Cavent/backend/config"
	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"github.com/HappyNaCl/Cavent/backend/infrastructure/persistence"
	"github.com/golang-jwt/jwt"
	"github.com/markbates/goth"
)

type JWTClaims struct {
	ProviderId string `json:"providerId"`
	Provider string `json:"provider"`
	Name string `json:"name"`
	Email string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
	Exp int `json:"exp"`
}

func GenerateJWT(claim JWTClaims) (string, error){
	claims := jwt.MapClaims{
		"providerId" : claim.ProviderId,
		"provider" : claim.Provider,
		"email" : claim.Email,
		"name": claim.Name,
		"avatarUrl" : claim.AvatarUrl,
		"exp" : time.Now().Add(time.Duration(claim.Exp) * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	byteSecret := []byte(jwtSecret)
	return token.SignedString(byteSecret)
}

func RegisterOrLoginOauthUser(user goth.User, provider string) (*model.User, error){
	return persistence.UserRepository(config.Database).RegisterOrLoginOauthUser(user, provider)
}

func LoginUser(email string, password string) (*model.User, error){
	user, err := persistence.UserRepository(config.Database).FindByEmail(email)
	if err != nil {
		return nil, err
	}
	log.Println(password)
	log.Println(user.Password)
	result, err := config.ComparePasswordAndHash(password, user.Password)

	if err != nil {
		return nil, err
	}

	if !result {
		return nil, errors.New("invalid password")
	}

	return user, nil
}