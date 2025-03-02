package routes

import (
	"order-management/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes routes
func SetupRoutes(router *gin.Engine, orderController *controllers.OrderController) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("/", orderController.CreateOrder)
		orderRoutes.GET("/:id", orderController.GetOrderStatus)
		orderRoutes.PUT("/:id", orderController.UpdateOrderStatus)
		router.GET("/orders", orderController.GetAllOrders)
	}
}
