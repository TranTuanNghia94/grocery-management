package services

import (
	"grocery-management/internal/config"
	"grocery-management/internal/models"
	repo "grocery-management/internal/repositories/user"
	"grocery-management/pkg/wrapper"

	"github.com/gin-gonic/gin"
)

// Ensure UserService implements IUserService
var _ IUserService = (*UserService)(nil)

// Define IUserService interface
type IUserService interface {
	GetUserByUsername(ctx *gin.Context, username string) (*models.User, *wrapper.ErrorWrapper)
	Login(ctx *gin.Context, username, password string) (*string, *wrapper.ErrorWrapper)
}

type UserService struct {
	repo   repo.IUserRepository
	config *config.Config
}

func NewUserService(repo repo.IUserRepository, config *config.Config) IUserService {
	return &UserService{
		repo:   repo,
		config: config,
	}
}
