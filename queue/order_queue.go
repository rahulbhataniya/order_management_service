package queue

import (
	"fmt"
	"log"
	"time"
)

// OrderQueue holds pending orders
var OrderQueue = make(chan int64, 100)

// ProcessOrders listens for new orders and processes them
func ProcessOrders(updateOrderStatus func(orderID int64, status string) error) {
	for orderID := range OrderQueue {
		fmt.Printf("Processing order: %d\n", orderID)

		// Simulate order processing time
		time.Sleep(2 * time.Second)

		// Update order status to "Completed"
		err := updateOrderStatus(orderID, "Completed")
		if err != nil {
			log.Printf("Failed to update order %d: %v\n", orderID, err)
		} else {
			fmt.Printf("Order %d processed successfully.\n", orderID)
		}
	}
}
