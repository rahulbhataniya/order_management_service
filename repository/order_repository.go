package repository

import (
	"order-management/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

// NewOrderRepository creates a new repository instance
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// GetOrderByID fetches an order by ID
func (r *OrderRepository) GetOrderByID(orderID int64) (*models.Order, error) {
	var order models.Order
	err := r.DB.First(&order, orderID).Error
	return &order, err
}

// UpdateOrderStatus updates the status of an order
func (r *OrderRepository) UpdateOrderStatus(orderID int64, status string) error {
	return r.DB.Model(&models.Order{}).Where("order_id = ?", orderID).Update("status", status).Error
}


// GetAllOrders fetches all orders from the database
func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	result := r.DB.Find(&orders) // GORM fetches all records
	return orders, result.Error
}

// CreateOrder saves an order to the database with default status
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	// Ensure the order status is set to "Pending"
	if order.Status == "" {
		order.Status = "Pending"
	}
	return r.DB.Create(order).Error
}

// GetOrderStatusCount returns the count of orders grouped by status
func (r *OrderRepository) GetOrderStatusCount() (map[string]int64, error) {
	var result []struct {
		Status string
		Count  int64
	}
	statusCounts := make(map[string]int64)

	err := r.DB.Table("orders").Select("status, COUNT(*) as count").Group("status").Scan(&result).Error
	if err != nil {
		return nil, err
	}

	for _, row := range result {
		statusCounts[row.Status] = row.Count
	}
	return statusCounts, nil
}