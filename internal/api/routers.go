package api

import (
	"grocery-management/internal/di"
	"grocery-management/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(container *di.Container) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())

	// Public routes
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  200,
		})
	})

	r.POST("/api/groceries", container.GroceryCtrl.GetAllGrocery)

	// r.POST("/api/login", container.UserController.Login)

	// Protected routes
	// api := r.Group("/api")
	// api.Use(middlewares.AuthMiddleware(container.Config))
	// {
	// 	api.GET("/users", container.UserController.ListUsers)
	// 	api.GET("/users/:id", container.UserController.GetUser)
	// 	api.PUT("/users/:id", container.UserController.UpdateUser)
	// 	api.DELETE("/users/:id", container.UserController.DeleteUser)
	// }

	return r
}
