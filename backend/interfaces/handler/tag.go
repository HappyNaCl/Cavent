package handler

import (
	"net/http"

	"github.com/HappyNaCl/Cavent/backend/config"
	"github.com/HappyNaCl/Cavent/backend/infrastructure/persistence"
	"github.com/HappyNaCl/Cavent/backend/interfaces/dto"
	"github.com/gin-gonic/gin"
)

type TagHandler struct {}

func NewTagHandler() *TagHandler {
	return &TagHandler{}
}

// GetAllTagsWithType godoc
// @Summary Get all tags with type
// @Description Get all tags with type
// @Tags tags
// @Accept json
// @Produce json
// @Success 200 {object} dto.TagTypeDTO
// @Failure 500 {object} responses.ErrorResponse
// @Router /tags [get]
// @Security ApiKeyAuth
func (h TagHandler) GetAllTagsWithType(c *gin.Context){
	tags, err := persistence.TagRepository(config.Database).GetAllTagsWithType()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var tagTypeDTOs []dto.TagTypeDTO
	for _, tt := range tags {
		tagType := dto.TagTypeDTO{
			Id:   tt.Id,
			Name: tt.Name,
			Tags: []dto.TagDTO{},
		}
	
		for _, tag := range tt.Tags {
			tagType.Tags = append(tagType.Tags, dto.TagDTO{
				Id:   tag.Id,
				Name: tag.Name,
			})
		}
	
		tagTypeDTOs = append(tagTypeDTOs, tagType)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": tagTypeDTOs,
	})
}