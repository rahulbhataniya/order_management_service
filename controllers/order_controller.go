package controllers

import (
	"net/http"
	"strconv"

	"order-management/models"
	"order-management/services"
	"order-management/queue"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService *services.OrderService
}

// NewOrderController initializes controller
func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

// Queue instance
var OrderQueue *queue.OrderQueue

// CreateOrderHandler handles order creation
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.OrderService.CreateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	OrderQueue.AddOrder(order.OrderID) // Add order to queue
	ctx.JSON(http.StatusCreated, gin.H{"message": "Order created", "order": order})
}

// GetOrder retrieves order by ID
func (c *OrderController) GetOrder(ctx *gin.Context) {
	orderID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := c.OrderService.GetOrderByID(orderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// GetAllOrders fetches all orders
func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.OrderService.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

// GetOrderStatus fetches the status of a specific order
func (c *OrderController) GetOrderStatus(ctx *gin.Context) {
	orderID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := c.OrderService.GetOrderByID(orderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"order_id": order.OrderID, "status": order.Status})
}

// UpdateOrderStatus updates the status of a specific order
func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
	orderID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var updateData struct {
		Status string `json:"status" binding:"required"`
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


// GetOrderMetricsHandler returns order processing metrics
func (c *OrderController) GetOrderMetricsHandler(ctx *gin.Context) {
	metrics, err := c.OrderService.GetOrderMetrics()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch metrics"})
		return
	}
	ctx.JSON(http.StatusOK, metrics)
}


// GetOrderDetails - Returns full order details
func (c *OrderController) GetOrderDetails(ctx *gin.Context) {
	orderID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := c.OrderService.GetOrderDetails(orderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

