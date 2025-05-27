package v1

import (
	"net/http"
	"os"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		jwtSecret := []byte(os.Getenv("ACCESS_JWT_SECRET"))
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
				Error: "unauthorized",
			})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
			return jwtSecret, nil
		})

		if err != nil  || !token.Valid {
			c.JSON(http.StatusUnauthorized,  &types.ErrorResponse{
				Error: "unauthorized",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized,  &types.ErrorResponse{
				Error: "invalid claims",
			})
			c.Abort()
			return
		}

		expFloat, ok := claims["exp"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
				Error: "invalid exp claim",
			})
			c.Abort()
			return
		}
		exp := int64(expFloat)
		now := time.Now().Unix()

		if now > exp {
			c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
				Error: "token expired",
			})
			c.Abort()
			return
		}

		for _, key := range []string{"sub", "exp", "iat", "email", "firstTimeLogin", "role"} {
			if val, ok := claims[key]; ok {
				c.Set(key, val)
			}
		}

		c.Next()
	}
}

func UnauthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString := c.Request.Header.Get("Authorization")
		tokenString = tokenString[len("Bearer "):]
		
		if tokenString != "" {
			c.JSON(http.StatusConflict,  &types.ErrorResponse{
				Error: "already logged in",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}