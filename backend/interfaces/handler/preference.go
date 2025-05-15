package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/HappyNaCl/Cavent/backend/application"
	"github.com/gin-gonic/gin"
)

type PreferenceHandler struct {}

func NewPreferenceHandler() *PreferenceHandler {
	return &PreferenceHandler{}
}

// UpdatePreferences godoc
// @Summary Update user preferences
// @Description Update user preferences
// @Tags user
// @Accept json
// @Produce json
// @Param userId formData string true "User ID"
// @Param preferences formData []string true "Preferences"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /user/preference [put]
// @Security ApiKeyAuth
func (h PreferenceHandler) UpdatePreferences(c *gin.Context) {
	userId := c.PostForm("userId")
	preferencesJson := c.PostForm("preferences")

	var preference []string

	err := json.Unmarshal([]byte(preferencesJson), &preference)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err = application.UpdatePrefences(userId, preference)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Preferences updated successfully")
	if firstTimeLogin, exists := c.Get("firstTimeLogin"); exists {
		log.Println("firstTimeLogin:", firstTimeLogin)
		if firstTimeLogin == true {
			id, _ := c.Get("id")
	
			provider, _ := c.Get("provider")
			
			email, _ := c.Get("email")
		
			avatarUrl, _ := c.Get("avatarUrl")
		
			name, _ := c.Get("name")
			log.Println("firstTimeLogin:", firstTimeLogin)
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