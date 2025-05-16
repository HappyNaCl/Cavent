package types

type JWTClaims struct {
	Sub           	string `json:"sub"`
	Exp           	int64  `json:"exp"`
	Iat           	int64  `json:"iat"`
	Email         	string `json:"email"`
	FirstTimeLogin 	bool   `json:"firstTimeLogin"`
	Role          	string `json:"role"`
}