package services

import (
	"order-management/models"
	"order-management/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
}

// NewOrderService initializes service
func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

// CreateOrder adds an order
func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.OrderRepo.CreateOrder(order)
}

// GetOrderByID retrieves an order by ID
func (s *OrderService) GetOrderByID(orderID int64) (*models.Order, error) {
	return s.OrderRepo.GetOrderByID(orderID)
}

// GetAllOrders fetches all orders
func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.OrderRepo.GetAllOrders()
}

// GetOrderStatus fetches the status of an order
func (s *OrderService) GetOrderStatus(orderID int64) (string, error) {
	return s.OrderRepo.GetOrderStatus(orderID)
}

// UpdateOrderStatus updates an order's status
func (s *OrderService) UpdateOrderStatus(orderID int64, status string) error {
	return s.OrderRepo.UpdateOrderStatus(orderID, status)
}

// GetOrderStatusCount returns the count of orders with a specific status
func (s *OrderService) GetOrderStatusCount(status string) (int64, error) {
	return s.OrderRepo.GetOrderStatusCount(status)
}

