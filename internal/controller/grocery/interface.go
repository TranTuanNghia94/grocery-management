package controller

import (
	services "grocery-management/internal/service/grocery"

	"github.com/gin-gonic/gin"
)

type IGroceryController interface {
	CreateGrocery(ctx *gin.Context)
}

// Ensure UserController implements IUserController
var _ IGroceryController = (*GroceryController)(nil)

type GroceryController struct {
	service services.IGroceryService
}

func NewGroceryController(service services.IGroceryService) IGroceryController {
	return &GroceryController{service: service}
}
