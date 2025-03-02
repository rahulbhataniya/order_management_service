package main

import (
	"order-management/config"
	"order-management/controllers"
	"order-management/repository"
	"order-management/services"
	"order-management/queue"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db := config.ConnectDB()

	// Initialize repository, service, and controller
	orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	// ✅ Initialize the queue correctly
	controllers.OrderQueue = queue.NewOrderQueue(orderService)

	// Set up Gin router
	r := gin.Default()

	// Define routes
	orderRoutes := r.Group("/orders")
	{
		orderRoutes.POST("/", orderController.CreateOrder)                     // Create order
		orderRoutes.GET("/:id", orderController.GetOrderStatus)                // Get order by ID
		orderRoutes.PUT("/:id/status", orderController.UpdateOrderStatus)      // Update order status
		orderRoutes.GET("/", orderController.GetAllOrders)                     // Get all orders
		orderRoutes.GET("/status/count", orderController.GetOrderStatusCount)  // Get order status counts
	}

	// Start server
	r.Run(":8080")
}
