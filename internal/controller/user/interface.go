package controller

import (
	services "grocery-management/internal/service/user"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Login(ctx *gin.Context)
}

// Ensure UserController implements IUserController
var _ IUserController = (*UserController)(nil)

type UserController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) IUserController {
	return &UserController{service: service}
}
