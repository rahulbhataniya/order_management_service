package repository

import (
	"order-management/models"
	"gorm.io/gorm"
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

// GetOrderStatusCount fetches the count of orders with a specific status
func (r *OrderRepository) GetOrderStatusCount(status string) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Order{}).
		Where("status = ?", status).
		Count(&count).Error
	return count, err
}