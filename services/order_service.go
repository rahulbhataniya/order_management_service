package services

import (
	"order-management/models"
	"order-management/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
}

// NewOrderService initializes the service
func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

// CreateOrder processes and saves an order
func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.OrderRepo.CreateOrder(order)
}

// GetOrderStatus fetches the order status
func (s *OrderService) GetOrderStatus(orderID int64) (*models.Order, error) {
	return s.OrderRepo.GetOrderByID(orderID)
}

// UpdateOrderStatus changes the status of an order
func (s *OrderService) UpdateOrderStatus(orderID int64, status string) error {
	return s.OrderRepo.UpdateOrderStatus(orderID, status)
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.OrderRepo.GetAllOrders()
}