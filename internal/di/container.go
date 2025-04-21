package di

import (
	"grocery-management/internal/config"
	groceryCtrl "grocery-management/internal/controller/grocery"
	userCtrl "grocery-management/internal/controller/user"
	groceryRepo "grocery-management/internal/repositories/grocery"
	userRepo "grocery-management/internal/repositories/user"
	grocerySvc "grocery-management/internal/service/grocery"
	userSvc "grocery-management/internal/service/user"

	"gorm.io/gorm"
)

type Container struct {
	DB          *gorm.DB
	Config      *config.Config
	GroceryRepo groceryRepo.IGroceryRepository
	GrocerySvc  grocerySvc.IGroceryService
	GroceryCtrl groceryCtrl.IGroceryController

	UserRepo userRepo.IUserRepository
	UserSvc  userSvc.IUserService
	UserCtrl userCtrl.IUserController
}

// NewContainer creates a new dependency injection container
func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
	// Initialize repositories
	groceryRepo := groceryRepo.NewGroceryRepo(db)

	// Initialize services
	grocerySvc := grocerySvc.NewGroceryService(groceryRepo, cfg)

	// Initialize controllers
	groceryCrtl := groceryCtrl.NewGroceryController(grocerySvc)

	userRepo := userRepo.NewUserRepository(db)
	userSvc := userSvc.NewUserService(userRepo, cfg)
	userCrtl := userCtrl.NewUserController(userSvc)

	return &Container{
		DB:          db,
		Config:      cfg,
		GroceryRepo: groceryRepo,
		GrocerySvc:  grocerySvc,
		GroceryCtrl: groceryCrtl,
		UserRepo:    userRepo,
		UserSvc:     userSvc,
		UserCtrl:    userCrtl,
	}
}
