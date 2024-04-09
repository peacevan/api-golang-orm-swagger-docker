package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peacevan/api-golang-orm-swagger-docker/schemas"
)

// @BasePath /api/v1

// @Summary Delete order
// @Description Delete a new job order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id query string true "Order identification"
// @Success 200 {object} DeleteOrderResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /order [delete]
func DeleteOrderHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	order := schemas.Order{}
	// Find Order
	if err := db.First(&order, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("order with id: %s not found", id))
		return
	}
	// Delete Order
	if err := db.Delete(&order).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting order with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-order", order)
}
