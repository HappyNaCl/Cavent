package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/HappyNaCl/Cavent/backend/application"
	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"github.com/HappyNaCl/Cavent/backend/interfaces/dto"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {}

// LoginByCredential godoc
// @Summary Login with email and password
// @Description Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param rememberMe formData boolean false "Remember Me"
// @Success 200 {object} dto.UserAuthDto
// @Failure 400 {object} model.ErrorResponse
// @Router /auth/login [post]
func (h AuthHandler) LoginCredential(c *gin.Context){
	email := c.PostForm("email")
	password := c.PostForm("password")
	rememberMe := c.PostForm("rememberMe")	

	time := 3600*24

	if rememberMe == "true" {
		time = 3600*24*7
	} 

	user, err := application.LoginUser(email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credential"})
		return
	}

	token, err := application.GenerateJWT(application.JWTClaims{
		Id: user.Id,
		Provider: user.Provider,
		Email: user.Email,
		AvatarUrl: user.AvatarUrl,
		Name: user.Name,
		Exp: time,
		FirstTimeLogin: user.FirstTimeLogin,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	appDomain := os.Getenv("APP_DOMAIN")

	c.SetCookie("token", token, time, "/", appDomain, false, true)
	c.JSON(200, gin.H{
		"message": "success",
		"data": &dto.UserAuthDto{
			Provider: user.Provider,
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			AvatarUrl: user.AvatarUrl,
		},
	})
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user 
// @Tags auth
// @Accept json
// @Produce json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param name formData string true "Name"
// @Success 200 {object} dto.UserAuthDto
// @Failure 400 {object} model.ErrorResponse
// @Router /auth/register [post]
func (h AuthHandler) RegisterUser(c *gin.Context){
	name := c.PostForm("Name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := application.RegisterUser(name, email, password)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := application.GenerateJWT(application.JWTClaims{
		Id: user.Id,
		Provider: user.Provider,
		Email: user.Email,
		AvatarUrl: user.AvatarUrl,
		Name: user.Name,
		FirstTimeLogin: user.FirstTimeLogin,
		Exp: 3600*24,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appDomain := os.Getenv("APP_DOMAIN")
	c.SetCookie("token", token, 3600*24, "/", appDomain, false, true)

	c.JSON(200, gin.H{
		"message": "success",
		"data": &dto.UserAuthDto{
			Provider: user.Provider,
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			AvatarUrl: user.AvatarUrl,
		},
	})
}

func (h AuthHandler) LoginWithOAuth(c *gin.Context){	
	provider := c.Param("provider")
	fmt.Println(provider)
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (h AuthHandler) LoginWithOAuthCallback(c *gin.Context){
	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	user, err := application.RegisterOrLoginOauthUser(gothUser, gothUser.Provider)
	if err != nil || user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	token, err := application.GenerateJWT(application.JWTClaims{
		Id: user.Id,
		Provider: user.Provider,
		Email: user.Email,
		AvatarUrl: user.AvatarUrl,
		Name: user.Name,
		FirstTimeLogin: user.FirstTimeLogin,
		Exp: 3600*24,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	appDomain := os.Getenv("APP_DOMAIN")
	c.SetCookie("token", token, 3600*24, "/", appDomain, false, true)

	c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL"))
}

func (h AuthHandler) Logout(c *gin.Context){
	c.SetCookie("token", "", -1, "/", os.Getenv("APP_DOMAIN"), false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
	})
}

func (h AuthHandler) CheckMe(c *gin.Context){
	id, _ := c.Get("id")
	
	provider, _ := c.Get("provider")
	
	email, _ := c.Get("email")

	avatarUrl, _ := c.Get("avatarUrl")

	name, _ := c.Get("name")

	firstTimeLogin, _ := c.Get("firstTimeLogin")

	c.JSON(http.StatusOK, gin.H{"user": model.User{
		Id: id.(string),
		Provider: provider.(string),
		Email: email.(string),
		Name: name.(string),
		AvatarUrl: avatarUrl.(string),
		FirstTimeLogin: firstTimeLogin.(bool),
	}})
}