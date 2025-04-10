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

type ContextProvider string
const providerKey ContextProvider = "provider"

func Run(port int) error{
	r := setupRoutes()
	log.Printf("Server running on port %d", port)
	return r.Run(fmt.Sprintf(":%d", port));
}

func setupRoutes() *gin.Engine{
	r := gin.Default()

	r.GET("/", index)

	auth := r.Group("/api/auth")
	auth.Use(UnauthMiddleware())
	auth.GET("/:provider", loginWithOAuth)
	auth.GET("/:provider/callback", loginWithOAuthCallback)

	auth.Use(AuthMiddleware())
	auth.GET("/logout", logout)

	protected := r.Group("/api/protected")
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

func loginWithoutOAuth(c *gin.Context){
	
}

func loginWithOAuth(c *gin.Context){	
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), providerKey, provider))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func loginWithOAuthCallback(c *gin.Context){
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	token, err := application.GenerateJWT(user.Email, user.Provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	appDomain := os.Getenv("APP_DOMAIN")
	log.Println(appDomain)
	c.SetCookie("token", token, 3600*24, "/", appDomain, false, true)

	loginUser, err := application.RegisterOrLoginOauthUser(user, user.Provider)
	if err != nil || loginUser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"provider": loginUser.Provider,
		"providerId": loginUser.ProviderID,
		"name": loginUser.Name,
		"email": loginUser.Email,
		"avatar": loginUser.AvatarUrl,
	})
}

func logout(c *gin.Context){
	c.SetCookie("token", "", -1, "/", os.Getenv("APP_DOMAIN"), false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
	})
}