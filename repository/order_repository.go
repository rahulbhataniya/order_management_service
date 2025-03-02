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

// CreateOrder saves an order to the database
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return r.DB.Create(order).Error
}

// GetOrderByID fetches an order by ID
func (r *OrderRepository) GetOrderByID(orderID int64) (*models.Order, error) {
	var order models.Order
	err := r.DB.First(&order, orderID).Error
	return &order, err
}

// UpdateOrderStatus updates the order status
func (r *OrderRepository) UpdateOrderStatus(orderID int64, status string) error {
	return r.DB.Model(&models.Order{}).Where("order_id = ?", orderID).Update("status", status).Error
}

// GetAllOrders fetches all orders from the database
func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	result := r.DB.Find(&orders) // GORM fetches all records
	return orders, result.Error
}