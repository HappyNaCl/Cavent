package interfaces

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
			return jwtSecret, nil
		})

		if err != nil  || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("email", claims["email"])
			c.Set("provider", claims["provider"])
		}

		c.Next()
	}
}

func UnauthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		token, err := c.Cookie("token")
		if err == nil && token != "" {
			c.JSON(http.StatusConflict, gin.H{"error": "You are already logged in"})
			c.Abort()
			return
		}
		c.Next()
	}
}