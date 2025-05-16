package v1

import (
	"net/http"
	"os"

	"github.com/HappyNaCl/Cavent/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		tokenString := c.Request.Header.Get("Authorization")
		tokenString = tokenString[len("Bearer "):]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, &types.ErrorResponse{
				Error: "unauthorized",
			})
			c.Abort()
			return
		}

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
		token, err := c.Cookie("token")
		tokenString := token[len("Bearer "):]
		
		if err == nil && tokenString != "" {
			c.JSON(http.StatusConflict,  &types.ErrorResponse{
				Error: "already logged in",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}