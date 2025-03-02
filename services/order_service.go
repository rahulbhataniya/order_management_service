package services

import (
	"order-management/models"
	"order-management/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
}

// NewOrderService creates a new service instance
func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: repo}
}

// CreateOrder creates an order
func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.OrderRepo.CreateOrder(order)
}

// GetOrderStatus retrieves order by ID
func (s *OrderService) GetOrderStatus(orderID int64) (*models.Order, error) {
	return s.OrderRepo.GetOrderByID(orderID)
}

// UpdateOrderStatus updates the order status
func (s *OrderService) UpdateOrderStatus(orderID int64, status string) error {
	return s.OrderRepo.UpdateOrderStatus(orderID, status)
}

// GetAllOrders fetches all orders
func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.OrderRepo.GetAllOrders()
}

// GetOrderStatusCount fetches the count of orders for each status
func (s *OrderService) GetOrderStatusCount() (map[string]int64, error) {
	return s.OrderRepo.GetOrderStatusCount()
}
