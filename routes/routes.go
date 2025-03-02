package routes

import (
	"order-management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, orderController *controllers.OrderController) {
	api := router.Group("/api")
	{
		api.GET("/orders", orderController.GetAllOrders)
		api.GET("/orders/:id", orderController.GetOrderDetails)   // Get full order details
		api.GET("/orders/:id/status", orderController.GetOrderStatus) // Get only order status
		api.POST("/orders", orderController.CreateOrder)
		api.PUT("/orders/:id/status", orderController.UpdateOrderStatus)

		// New route for order metrics
		api.GET("/orders/metrics", orderController.GetOrderMetricsHandler)
	}
}

