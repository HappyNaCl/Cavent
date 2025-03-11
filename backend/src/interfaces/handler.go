package interfaces

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HappyNaCl/Cavent/src/application"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Run(port int) error{
	r := setupRoutes()
	log.Printf("Server running on port %d", port)
	return r.Run(fmt.Sprintf(":%d", port));
}

func setupRoutes() *gin.Engine{
	r := gin.Default()

	r.GET("/", index)
	r.GET("/auth/:provider", loginWithOAuth)
	r.GET("/auth/:provider/callback", loginWithOAuthCallback)

	protected := r.Group("/protected")
	protected.Use(AuthMiddleware())
	protected.GET("/", protectedIndex)
	return r
}

func index(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func protectedIndex(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Protected",
	})
}

func loginWithOAuth(c *gin.Context){
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func loginWithOAuthCallback(c *gin.Context){
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	// fmt.Println(user)

	token, err := application.GenerateJWT(user.Email, user.Provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	appDomain := os.Getenv("APP_DOMAIN")
	log.Println(appDomain)
	c.SetCookie("token", token, 3600*24, "/", appDomain, false, true)

	application.RegisterOrLoginUser(user, user.Provider)
}
