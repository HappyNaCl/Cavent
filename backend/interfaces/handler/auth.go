package handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/HappyNaCl/Cavent/backend/application"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {}

type ContextProvider string
const providerKey ContextProvider = "provider"

func (h AuthHandler) LoginCredential(c *gin.Context){
	email := c.PostForm("email")
	password := c.PostForm("password")
	user, err := application.LoginUser(email, password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := application.GenerateJWT(user.Email, user.Provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	appDomain := os.Getenv("APP_DOMAIN")

	c.SetCookie("token", token, 3600*24, "/", appDomain, false, true)
	c.JSON(200, gin.H{
		"message": "success",
		"data": gin.H{
			"provider": user.Provider,
			"providerId": user.ProviderID,
			"name": user.Name,
			"email": user.Email,
			"avatar": user.AvatarUrl,
		},
	})

}

func (h AuthHandler) RegisterUser(c *gin.Context){
	fullName := c.PostForm("fullName")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := application.RegisterUser(fullName, email, password)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := application.GenerateJWT(user.Email, user.Provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appDomain := os.Getenv("APP_DOMAIN")
	c.SetCookie("token", token, 3600*24, "/", appDomain, false, true)

	c.JSON(200, gin.H{
		"message": "success",
		"data": gin.H{
			"provider": user.Provider,
			"providerId": user.ProviderID,
			"name": user.Name,
			"email": user.Email,
			"avatar": user.AvatarUrl,
		},
	})
}

func (h AuthHandler) LoginWithOAuth(c *gin.Context){	
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), providerKey, provider))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (h AuthHandler) LoginWithOAuthCallback(c *gin.Context){
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

	c.JSON(200, gin.H{
		"message": "success",
		"data": gin.H{
			"provider": loginUser.Provider,
			"providerId": loginUser.ProviderID,
			"name": loginUser.Name,
			"email": user.Email,
			"avatar": loginUser.AvatarUrl,
		},
	})
}

func (h AuthHandler) Logout(c *gin.Context){
	c.SetCookie("token", "", -1, "/", os.Getenv("APP_DOMAIN"), false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
	})
}