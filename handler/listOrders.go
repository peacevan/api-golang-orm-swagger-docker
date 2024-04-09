package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peacevan/api-golang-orm-swagger-docker/schemas"
)

// @BasePath /api/v1

// @Summary List orders
// @Description List all job orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} ListOrdersResponse
// @Failure 500 {object} ErrorResponse
// @Router /orders [get]
func ListOrdersHandler(ctx *gin.Context) {
	orders := []schemas.Order{}

	if err := db.Find(&orders).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing orders")
		return
	}

	sendSuccess(ctx, "list-orders", orders)
}
