package handler

import (
	"net/http"

	"github.com/HappyNaCl/Cavent/backend/application"
	"github.com/HappyNaCl/Cavent/backend/interfaces/dto"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}


func (h UserHandler) GetUserTag(c *gin.Context) {
	userId, _ := c.Get("id")
	userIdStr, ok := userId.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tags, err := application.GetUserTag(userIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tagDto []dto.TagDTO

	for _, tag := range tags {
		tagDto = append(tagDto, dto.TagDTO{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": tagDto,
	})
}