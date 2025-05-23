package types

import (
	"errors"
	"time"
)

type JWTClaims struct {
	Sub           	string `json:"sub"`
	Exp           	int64  `json:"exp"`
	Iat           	int64  `json:"iat"`
	Email         	string `json:"email"`
	FirstTimeLogin 	bool   `json:"firstTimeLogin"`
	Role          	string `json:"role"`
}

func (c JWTClaims) Valid() error {
	now := time.Now().Unix()

	if c.Exp < now {
		return errors.New("token is expired")
	}
	if c.Iat > now {
		return errors.New("token used before issued")
	}

	return nil
}