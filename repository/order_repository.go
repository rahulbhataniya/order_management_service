package repository

import (
	"order-management/models"
	"gorm.io/gorm"
	"database/sql"
)

type OrderRepository struct {
	DB *gorm.DB
}

// NewOrderRepository initializes repository
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// CreateOrder inserts an order
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return r.DB.Create(order).Error
}

// GetOrderByID fetches an order by ID
func (r *OrderRepository) GetOrderByID(orderID int64) (*models.Order, error) {
	var order models.Order
	err := r.DB.Where("order_id = ?", orderID).First(&order).Error
	return &order, err
}

// GetAllOrders retrieves all orders
func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	result := r.DB.Find(&orders)
	return orders, result.Error
}

// UpdateMultipleOrdersStatus updates multiple orders in one DB transaction
func (r *OrderRepository) UpdateMultipleOrdersStatus(orderIDs []int64, status string) error {
	return r.DB.Model(&models.Order{}).
		Where("order_id IN ?", orderIDs).
		Update("status", status).
		Error
}

// GetOrderStatus fetches only the status of an order by ID
func (r *OrderRepository) GetOrderStatus(orderID int64) (string, error) {
	var status string
	err := r.DB.Model(&models.Order{}).
		Select("status").
		Where("order_id = ?", orderID).
		Scan(&status).Error
	return status, err
}

// UpdateOrderStatus updates the status of an order
func (r *OrderRepository) UpdateOrderStatus(orderID int64, status string) error {
	return r.DB.Model(&models.Order{}).
		Where("order_id = ?", orderID).
		Update("status", status).
		Error
}

// GetTotalOrders retrieves the total number of orders
func (r *OrderRepository) GetTotalOrders() (int64, error) {
	var count int64
	err := r.DB.Model(&models.Order{}).Count(&count).Error
	return count, err
}

// GetOrderStatusCount fetches count of orders by status
func (r *OrderRepository) GetOrderStatusCount() (map[string]int64, error) {
	var result []struct {
		Status string
		Count  int64
	}
	statusCounts := make(map[string]int64)

	err := r.DB.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	for _, row := range result {
		statusCounts[row.Status] = row.Count
	}
	return statusCounts, nil
}

// GetAvgProcessingTime calculates the average processing time for completed orders
func (r *OrderRepository) GetAvgProcessingTime() (float64, error) {
	var avgTime sql.NullFloat64

	err := r.DB.Raw("SELECT AVG(TIMESTAMPDIFF(SECOND, created_at, updated_at)) FROM orders WHERE status = 'completed'").Scan(&avgTime).Error
	if err != nil {
		return 0.0, err
	}

	// If avgTime is NULL, return 0.0
	if !avgTime.Valid {
		return 0.0, nil
	}

	return avgTime.Float64, nil
}
