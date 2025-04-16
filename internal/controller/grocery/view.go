package controller

import (
	"grocery-management/pkg/logger"
	"grocery-management/pkg/wrapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *GroceryController) GetAllGrocery(ctx *gin.Context) {
	logger.Info("Get all groceries")

	groceries, err := c.service.GetAllGroceries()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, wrapper.ServerErrorResponse(err, "Failed to get groceries"))
		return
	}

	// Use the wrapper.SuccessResponse correctly
	response := wrapper.SuccessResponse(groceries, "Get all groceries successfully")

	ctx.JSON(http.StatusOK, response)
}
