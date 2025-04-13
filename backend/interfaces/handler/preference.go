package handler

import (
	"net/http"
	"os"

	"github.com/HappyNaCl/Cavent/backend/application"
	"github.com/gin-gonic/gin"
)

type PreferenceHandler struct {}

func NewPreferenceHandler() *PreferenceHandler {
	return &PreferenceHandler{}
}

func (h PreferenceHandler) UpdatePreferences(c *gin.Context) {
	if firstTimeLogin, exists := c.Get("firstTimeLogin"); exists {
		if firstTimeLogin == true {
			id, _ := c.Get("id")
	
			provider, _ := c.Get("provider")
			
			email, _ := c.Get("email")
		
			avatarUrl, _ := c.Get("avatarUrl")
		
			name, _ := c.Get("name")

			token, err := application.GenerateJWT(application.JWTClaims{
				Id: id.(string),
				Provider: provider.(string),
				Email: email.(string),
				AvatarUrl: avatarUrl.(string),
				Name: name.(string),
				FirstTimeLogin: false,
				Exp: 3600*24,
			})

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			
			appDomain := os.Getenv("APP_DOMAIN")
			c.SetCookie("token", token, 3600*24, "/", appDomain, false, true)	
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Preferences updated successfully"})
}