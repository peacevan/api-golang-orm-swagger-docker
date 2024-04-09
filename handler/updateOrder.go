package handler

import (
	"net/http"

	"github.com/arthur404dev/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update order
// @Description Update a job order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id query string true "Order Identification"
// @Param order body UpdateOrderRequest true "Order data to Update"
// @Success 200 {object} UpdateOrderResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /order [put]
func UpdateOrderHandler(ctx *gin.Context) {
	request := UpdateOrderRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	order := schemas.Order{}

	if err := db.First(&order, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "order not found")
		return
	}
	// Update order
	if request.Role != "" {
		order.Role = request.Role
	}

	if request.Company != "" {
		order.Company = request.Company
	}

	if request.Location != "" {
		order.Location = request.Location
	}

	if request.Remote != nil {
		order.Remote = *request.Remote
	}

	if request.Link != "" {
		order.Link = request.Link
	}

	if request.Salary > 0 {
		order.Salary = request.Salary
	}
	// Save order
	if err := db.Save(&order).Error; err != nil {
		logger.Errorf("error updating order: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating order")
		return
	}
	sendSuccess(ctx, "update-order", order)
}
