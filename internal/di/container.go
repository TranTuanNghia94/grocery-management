package di

import (
	"grocery-management/internal/config"
	controller "grocery-management/internal/controller/grocery"
	repo "grocery-management/internal/repositories/grocery"
	services "grocery-management/internal/service/grocery"

	"gorm.io/gorm"
)

type Container struct {
	DB          *gorm.DB
	Config      *config.Config
	GroceryRepo repo.IGroceryRepository
	GrocerySvc  services.IGroceryService
	GroceryCtrl controller.IGroceryController
}

// NewContainer creates a new dependency injection container
func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
	// Initialize repositories
	groceryRepo := repo.NewGroceryRepo(db)

	// Initialize services
	grocerySvc := services.NewGroceryService(groceryRepo, cfg)

	// Initialize controllers
	groceryCrtl := controller.NewGroceryController(grocerySvc)

	return &Container{
		DB:          db,
		Config:      cfg,
		GroceryRepo: groceryRepo,
		GrocerySvc:  grocerySvc,
		GroceryCtrl: groceryCrtl,
	}
}
