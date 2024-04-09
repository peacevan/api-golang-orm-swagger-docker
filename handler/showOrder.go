package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peacevan/api-golang-orm-swagger-docker/schemas"
)

// @BasePath /api/v1

// @Summary Show order
// @Description Show a job order
// @Tags orders
// @Accept json
// @Produce json
// @Param id query string true "Order identification"
// @Success 200 {object} ShowOrderResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /order [get]
func ShowOrderHandler(ctx *gin.Context) {
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
	sendSuccess(ctx, "show-order", order)
}
