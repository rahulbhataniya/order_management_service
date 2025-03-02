Order Management System
A scalable backend service for handling e-commerce orders, queuing, and real-time order status tracking.

Features

1. RESTful APIs for order management
2. In-memory queue for asynchronous processing
3. Scalable architecture to handle high load (1,000 concurrent orders)
4. Real-time order status tracking
5. Metrics API for reporting order processing time

Tech Stack:
Golang (Gin Framework)
MySQL
GORM (ORM for DB operations)
In-Memory Queue (For async processing)

API Endpoints:

1. Create an Order: Endpoint: POST /orders

http://localhost:8080/api/orders

Request(json):
{
"user_id": 101,
"item_ids": "5,6,7",
"total_amount": 299.99
}

Response:
{
"message": "Order created",
"order": {
"order_id": 3,
"user_id": 101,
"item_ids": "5,6,7",
"total_amount": 299.99,
"status": "pending",
"CreatedAt": 1740890095,
"UpdatedAt": 1740890095
}
}

2. Get Order by ID
   Endpoint: GET /orders/{id}

Postman:
http://localhost:8080/api/orders/3

Response: {
"order_id": 3,
"user_id": 101,
"item_ids": "5,6,7",
"total_amount": 299.99,
"status": "processing",
"CreatedAt": 1740890095,
"UpdatedAt": 1740890097
}

3. Get Order status by ID
   Endpoint: GET /orders/{id}/status

Postmen: http://localhost:8080/api/orders/3/status
Response:
{
"order_id": 3,
"status": "processing"
}

4. Update Order Status
   Endpoint: PUT /orders/{id}/status

Request:
{
"status": "processing"
}

Response:
{ "message": "Order status updated successfully" }

5. Get All Orders

Endpoint: GET /orders
Response:
[
{ "order_id": 1, "status": "processing" },
{ "order_id": 2, "status": "pending" }
]

6. Get Order Processing Metrics:
   Endpoint: GET /orders/metrics

Response:
{
"total_orders": 3,
"avg_processing_time": 0,
"pending_orders": 0,
"processing_orders": 1,
"completed_orders": 0
}
