package queue

import (
	"fmt"
	"order-management/services"
	"sync"
	"time"
)

type OrderQueue struct {
	queue       chan int64
	orderService *services.OrderService
	wg          sync.WaitGroup
}

// NewOrderQueue initializes a queue
func NewOrderQueue(orderService *services.OrderService) *OrderQueue {
	q := &OrderQueue{
		queue:       make(chan int64, 1000), // Buffer size 1000 for high-load processing
		orderService: orderService,
	}
	go q.ProcessQueue()
	return q
}

// AddOrder adds an order ID to the queue
func (q *OrderQueue) AddOrder(orderID int64) {
	q.queue <- orderID
}

// ProcessQueue handles queued orders
func (q *OrderQueue) ProcessQueue() {
	for orderID := range q.queue {
		q.wg.Add(1)
		go q.processOrder(orderID)
	}
}

// processOrder processes each order
func (q *OrderQueue) processOrder(orderID int64) {
	defer q.wg.Done()
	time.Sleep(2 * time.Second) // Simulating processing delay

	err := q.orderService.UpdateOrderStatus(orderID, "processing")
	if err != nil {
		fmt.Println("Error updating order status:", err)
		return
	}

	fmt.Printf("Order %d processed successfully\n", orderID)
}
