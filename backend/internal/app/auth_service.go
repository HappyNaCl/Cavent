package app

import (
	"net/http"
	"regexp"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/postgresdb"
	"github.com/HappyNaCl/Cavent/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthService struct {
	authRepo repo.AuthRepository
	logger *zap.Logger
}

type RegisterRequest struct {
	Name 			string `json:"name"`
	Password 		string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Email 			string `json:"email"`
}

const (
	ErrInvalidEmail = "Invalid email"
	ErrNameLength = "Name must be 4 to 24 characters"
	ErrPasswordLength = "Password must be 8 to 24 characters"
	ErrInvalidPassword = "Password must contain a uppercase letter and a number"
	ErrConfirmPasswordMismatch = "Password and Confirm Password is not the same"
)

func NewAuthService(db *gorm.DB, redis *redis.Client, logger *zap.Logger) *AuthService {
	return &AuthService{
		authRepo: postgresdb.NewAuthGormRepo(db, redis, logger),
		logger: logger,
	}
}

func (as *AuthService) RegisterUser(c *gin.Context) {
	var req *RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: "invalid form requests",
		})
		c.Abort()
		return
	}

	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(req.Email) {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: ErrInvalidEmail,
		})
		return
	}

	if len(req.Name) < 4 || len(req.Name) > 24 {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: ErrNameLength,
		})
		return
	}

	if len(req.Password) < 8 || len(req.Password) > 32 {
        c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: ErrPasswordLength,
		})
		return
    }

	if !isValidPassword(req.Password) {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: ErrInvalidPassword,
		})
		return
	}
	
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			Error: ErrConfirmPasswordMismatch,
		})
		return
	}

	userFactory := factory.NewUserFactory()

	userModel := userFactory.GetUser(req.Name, req.Email, req.Password)

	_, err := as.authRepo.RegisterUser(userModel)
	if err != nil {
		c.JSON(http.StatusCreated, &types.SuccessResponse{
			Message: "Success",
			// Data: user,
		})
	}
}

func isValidPassword(password string) bool {
    hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
    hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

    return hasUppercase && hasNumber
}
