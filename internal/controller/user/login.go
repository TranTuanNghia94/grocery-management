package controller

import (
	"grocery-management/pkg/wrapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *UserController) Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, wrapper.ServerErrorResponse(err, "Invalid input"))
		return
	}

	jwt, err := c.service.Login(input.Username, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, wrapper.ServerErrorResponse(err, "Login failed"))
		return
	}

	response := wrapper.SuccessResponse(gin.H{"jwt": jwt}, "Login successfully")

	ctx.JSON(http.StatusOK, response)
}
