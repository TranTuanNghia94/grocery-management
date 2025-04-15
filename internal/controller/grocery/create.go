package controller

import (
	"grocery-management/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *GroceryController) CreateGrocery(ctx *gin.Context) {
	var input models.Grocery

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g, err := c.service.CreateGrocery(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": g})
}
