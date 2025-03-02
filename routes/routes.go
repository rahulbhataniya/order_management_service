package routes

import (
	"order-management/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes routes
func SetupRoutes(router *gin.Engine, orderController *controllers.OrderController) {
	// Define routes
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("/", orderController.CreateOrder)                     // Create order
		orderRoutes.GET("/:id", orderController.GetOrderStatus)                // Get order by ID
		orderRoutes.PUT("/:id/status", orderController.UpdateOrderStatus)      // Update order status
		orderRoutes.GET("/", orderController.GetAllOrders)                     // Get all orders
		orderRoutes.GET("/status/count", orderController.GetOrderStatusCount)  // Get order status counts
	}
}
