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

// GetAllCategory godoc
// @Summary      Get all categories
// @Description  Retrieve all categories
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param	     Authorization  header  string  true  "Bearer token for authentication"
// @Success      200  {object}  types.SuccessResponse{data=[]common.CategoryTypeResult}
// @Failure      500  {object}  types.ErrorResponse
// @Router       /category [get]
// @Security     BearerAuth
func (h *CategoryHandler) getAllCategory(c *gin.Context) {
	result, err := h.categoryService.GetAllCategory()
	if err != nil {
		c.JSON(500, types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &types.SuccessResponse{
		Message: "success",
		Data: result.CategoryTypes,
	})
}