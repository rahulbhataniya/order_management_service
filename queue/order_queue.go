package queue

import (
	"fmt"
	"order-management/services"
	"sync"
)

// OrderQueue represents an in-memory queue
type OrderQueue struct {
	orders    chan int64
	service   *services.OrderService
	waitGroup sync.WaitGroup
}

// NewOrderQueue initializes a new queue for processing orders
func NewOrderQueue(service *services.OrderService) *OrderQueue {
	queue := &OrderQueue{
		orders:  make(chan int64, 100), // Buffer size of 100
		service: service,
	}
	go queue.processOrders()
	return queue
}

// AddOrder adds an order ID to the queue
func (q *OrderQueue) AddOrder(orderID int64) {
	q.orders <- orderID
}

// processOrders continuously processes orders from the queue
func (q *OrderQueue) processOrders() {
	for orderID := range q.orders {
		fmt.Println("Processing order:", orderID)

		// Example: Update status to "Processing"
		q.service.UpdateOrderStatus(orderID, "Processing")

		// Simulate completion (in a real scenario, add business logic)
		q.service.UpdateOrderStatus(orderID, "Completed")

		fmt.Println("Order", orderID, "completed")
	}
}
