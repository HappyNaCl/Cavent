package service

import (
	e "errors"
	"regexp"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/mapper"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	authRepo repo.AuthRepository
	campusRepo repo.CampusRepository
}

func NewAuthService(db *gorm.DB, redis *redis.Client) *AuthService {
	return &AuthService{
		authRepo: postgresdb.NewAuthGormRepo(db),
		campusRepo: postgresdb.NewCampusGormRepo(db),
	}
}

func (as *AuthService) RegisterUser(com *command.RegisterUserCommand) (*command.RegisterUserCommandResult, error) {
	if !isValidEmail(com.Email) {
		return nil, errors.ErrInvalidEmail
	}

	if len(com.Name) < 4 || len(com.Name) > 24 {
		return nil, errors.ErrNameLength
	}

	if len(com.Password) < 8 || len(com.Password) > 32 {
		return nil, errors.ErrPasswordLength
    }

	if !isValidPassword(com.Password) {
		return nil, errors.ErrInvalidPassword
	}
	
	if com.Password != com.ConfirmPassword {
		return nil, errors.ErrConfirmPasswordMismatch
	}

	if com.InviteCode != nil && len(*com.InviteCode) != 6 {
		return nil, errors.ErrInviteCodeLength
	}

	userFactory := factory.NewUserFactory()
	userModel := userFactory.GetUser(com.Name, com.Email, com.Password)
	if com.InviteCode != nil {
		campus, err := as.campusRepo.FindCampusByInviteCode(*com.InviteCode)
		if err != nil {
			return nil, err
		}
		userModel.CampusId = &campus.Id
	} 

	user, err := as.authRepo.RegisterUser(userModel)

	if err != nil {
		var pgErr *pgconn.PgError
		if e.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, errors.ErrDuplicateEmail
		}
		return nil, err
	}

	userResult := mapper.NewUserResultFromRegisteredUser(user)

	return &command.RegisterUserCommandResult{
		Result: userResult,
	}, nil
}

func (as *AuthService) LoginUser(com *command.LoginUserCommand) (*command.LoginUserCommandResult, error) {
	user, err := as.authRepo.LoginUser(com.Email)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrInvalidCredentials
	}

	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(com.Password)) != nil {
		return nil, errors.ErrInvalidCredentials
	}

	userResult := mapper.NewUserResultFromLoginUser(user)
	
	return &command.LoginUserCommandResult{
		Result: userResult,
	}, nil
}

func (as *AuthService) HandleOAuth(com *command.HandleOAuthCommand) (*command.HandleOAuthCommandResult, error) {
	user, err := as.authRepo.RegisterOrLoginOauthUser(com.User, com.Provider)

	if err != nil {
		return nil, err
	}
	
	return &command.HandleOAuthCommandResult{
		Result: mapper.NewUserResultFromUser(user),
	}, nil
}

func isValidPassword(password string) bool {
    hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
    hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

    return hasUppercase && hasNumber
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}