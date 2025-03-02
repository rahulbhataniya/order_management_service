package main

import (
	"order-management/config"
	"order-management/controllers"
	"order-management/repository"
	"order-management/routes"
	"order-management/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	db := config.ConnectDB()

	// Setup Repository
	orderRepo := repository.NewOrderRepository(db)

	// Setup Service
	orderService := services.NewOrderService(orderRepo)

	// Setup Controller
	orderController := controllers.NewOrderController(orderService)

	// Initialize Gin Router
	router := gin.Default()
	routes.SetupRoutes(router, orderController)

	// Start Server
	router.Run(":8080")
}
