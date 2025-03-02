package main

import (
	"order-management/config"
	"order-management/controllers"
	"order-management/repository"
	"order-management/services"
	"order-management/queue"
	"order-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db := config.ConnectDB()

	// Initialize repository, service, and controller
	orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	// Initialize in-memory queue for async processing
	controllers.OrderQueue = queue.NewOrderQueue(orderService)

	router := gin.Default()
	routes.SetupRoutes(router, orderController)

	// Start server
	router.Run(":8080")
}
