# Order Management System

A scalable backend service for handling e-commerce orders with queuing and real-time order status tracking.

## Features

- RESTful APIs for seamless order management
- In-memory queue for asynchronous order processing
- Scalable architecture capable of handling high loads (1,000 concurrent orders)
- Real-time order status tracking for better customer experience
- Metrics API to monitor order processing performance

## Tech Stack

- Golang (Gin Framework) - Web framework for handling API requests
- MySQL - Relational database for order storage
- GORM - ORM for simplified database operations
- In-Memory Queue - For asynchronous order processing

## API Endpoints

### 1. Create an Order
Endpoint: `POST /orders`

URL: `http://localhost:8080/api/orders`

#### Request (JSON):
json
{
  "user_id": 101,
  "item_ids": "5,6,7",
  "total_amount": 299.99
}


#### Response:
json
{
  "message": "Order created",
  "order": {
    "order_id": 3,
    "user_id": 101,
    "item_ids": "5,6,7",
    "total_amount": 299.99,
    "status": "pending",
    "created_at": 1740890095,
    "updated_at": 1740890095
  }
}


---

### 2. Get Order by ID
Endpoint: `GET /orders/{id}`

URL: `http://localhost:8080/api/orders/3`

#### Response:
json
{
  "order_id": 3,
  "user_id": 101,
  "item_ids": "5,6,7",
  "total_amount": 299.99,
  "status": "processing",
  "created_at": 1740890095,
  "updated_at": 1740890097
}


---

### 3. Get Order Status by ID
Endpoint: `GET /orders/{id}/status`

URL: `http://localhost:8080/api/orders/3/status`

#### Response:
json
{
  "order_id": 3,
  "status": "processing"
}


---

### 4. Update Order Status
Endpoint: `PUT /orders/{id}/status`

URL: `http://localhost:8080/api/orders/3/status`

#### Request:
json
{
  "status": "processing"
}


#### Response:
json
{
  "message": "Order status updated successfully"
}


---

### 5. Get All Orders
Endpoint: `GET /orders`

URL: `http://localhost:8080/api/orders`

#### Response:
json
[
  { "order_id": 1, "status": "processing" },
  { "order_id": 2, "status": "pending" }
]


---

### 6. Get Order Processing Metrics
Endpoint: `GET /orders/metrics`

URL: `http://localhost:8080/api/orders/metrics`

#### Response:
json
{
  "total_orders": 3,
  "avg_processing_time": 0,
  "pending_orders": 0,
  "processing_orders": 1,
  "completed_orders": 0
}


## Public API URLs

For testing on a deployed environment, use the following endpoints:

- Create/Get Orders:  
  `http://ec2-65-0-32-133.ap-south-1.compute.amazonaws.com/api/orders`

- Get Order by ID:  
  `http://ec2-65-0-32-133.ap-south-1.compute.amazonaws.com/api/orders/1`

- Get Order Metrics:  
  `http://ec2-65-0-32-133.ap-south-1.compute.amazonaws.com/api/orders/metrics`

---

## Setup and Installation

1. Clone the repository:
   
   git clone <repository-url>
   cd order-management-system
   
2. Install dependencies:
   
   go mod tidy
   
3. Configure environment variables (Database connection, API keys, etc.).
4. Run the application:
   
   go run cmd/main.go
