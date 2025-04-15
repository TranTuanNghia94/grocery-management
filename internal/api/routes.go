package api

import (
	"grocery-management/internal/config"
	controller "grocery-management/internal/controller/grocery"
	repositories "grocery-management/internal/repositories/grocery"
	services "grocery-management/internal/service/grocery"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, log *zap.Logger) *gin.Engine {
	r := gin.Default()

	config, _ := config.LoadConfig(log)

	// Setup dependencies
	groceryRepo := repositories.NewUserRepository(db)
	groceryService := services.NewGroceryService(groceryRepo, config)
	groceryCotnroller := controller.NewGroceryController(groceryService)

	// Public routes
	r.POST("/api/register", groceryCotnroller.CreateGrocery)

	// Protected routes
	// api := r.Group("/api")
	// api.Use(middlewares.AuthMiddleware(config))
	// {
	// 	api.GET("/users", groceryCotnroller.CreateGrocery)
	// 	api.GET("/users/:id", userController.GetUser)
	// 	api.PUT("/users/:id", userController.UpdateUser)
	// 	api.DELETE("/users/:id", userController.DeleteUser)
	// }

	return r
}
