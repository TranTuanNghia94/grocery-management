package services

import (
	"grocery-management/internal/config"
	"grocery-management/internal/models"
	repo "grocery-management/internal/repositories/user"
)

// Ensure UserService implements IUserService
var _ IUserService = (*UserService)(nil)

// Define IUserService interface
type IUserService interface {
	GetUserByUsername(username string) (*models.User, error)
	Login(username, password string) (*models.User, error)
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
