package services

import (
	"grocery-management/internal/config"
	"grocery-management/internal/models"
	repo "grocery-management/internal/repositories/grocery"
)

// Ensure UserService implements IUserService
var _ IGroceryService = (*GroceryService)(nil)

type IGroceryService interface {
	CreateGrocery(input models.Grocery) (*models.Grocery, error)
}

type GroceryService struct {
	repo   repo.IGroceryRepository
	config *config.Config
}

func NewGroceryService(repo repo.IGroceryRepository, config *config.Config) IGroceryService {
	return &GroceryService{
		repo:   repo,
		config: config,
	}
}
