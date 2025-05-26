package v1

import (
	"net/http"

	"github.com/HappyNaCl/Cavent/backend/internal/app/service"
	"github.com/HappyNaCl/Cavent/backend/internal/interfaces/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func (h *CategoryHandler) SetupRoutes(v1 *gin.RouterGroup) {
	v1.GET("/category", h.GetAllCategory)
}

func NewCategoryHandler(db *gorm.DB, redis *redis.Client) types.Route {
	return &CategoryHandler{
		categoryService: service.NewCategoryService(db, redis),
	}
}

func (h *CategoryHandler) GetAllCategory(c *gin.Context) {
	result, err := h.categoryService.GetAllCategory()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: result.CategoryTypes,
	})
}
