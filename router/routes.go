package router

import (
	docs "github.com/arthur404dev/gopportunities/docs"
	"github.com/arthur404dev/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/order", handler.ShowOrderHandler)
		v1.POST("/order", handler.CreateOrderHandler)
		v1.DELETE("/order", handler.DeleteOrderHandler)
		v1.PUT("/order", handler.UpdateOrderHandler)
		v1.GET("/orders", handler.ListOrdersHandler)

	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
