package handler

import (
	"net/http"

	"github.com/arthur404dev/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create order
// @Description Create a new job order
// @Tags orders
// @Accept json
// @Produce json
// @Param request body CreateOrderRequest true "Request body"
// @Success 200 {object} CreateOrderResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /order [post]
func CreateOrderHandler(ctx *gin.Context) {
	request := CreateOrderRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	order := schemas.Order{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&order).Error; err != nil {
		logger.Errorf("error creating order: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating order on database")
		return
	}

	sendSuccess(ctx, "create-order", order)
}
