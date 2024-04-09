package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peacevan/api-golang-orm-swagger-docker/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOrderResponse struct {
	Message string                `json:"message"`
	Data    schemas.OrderResponse `json:"data"`
}

type DeleteOrderResponse struct {
	Message string                `json:"message"`
	Data    schemas.OrderResponse `json:"data"`
}
type ShowOrderResponse struct {
	Message string                `json:"message"`
	Data    schemas.OrderResponse `json:"data"`
}
type ListOrdersResponse struct {
	Message string                  `json:"message"`
	Data    []schemas.OrderResponse `json:"data"`
}
type UpdateOrderResponse struct {
	Message string                `json:"message"`
	Data    schemas.OrderResponse `json:"data"`
}
