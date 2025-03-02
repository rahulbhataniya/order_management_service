package controllers

import (
	"net/http"
	"strconv"

	"order-management/models"
	"order-management/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService *services.OrderService
}

// NewOrderController initializes the controller
func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

// CreateOrder handles order creation requests
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order

	// Bind JSON request body to order model
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save order using service
	if err := c.OrderService.CreateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

// GetOrderStatus fetches order details
func (c *OrderController) GetOrderStatus(ctx *gin.Context) {
	orderID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := c.OrderService.GetOrderStatus(orderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"order": order})
}

// UpdateOrderStatus updates the status of an order
func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
	orderID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var updateData struct {
		Status string `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.OrderService.UpdateOrderStatus(orderID, updateData.Status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}


func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.OrderService.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}