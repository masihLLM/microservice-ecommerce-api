# Ecommerce API

This project is a microservice-based REST API for an e-commerce platform, built using Go and the Gin framework. It consists of three main services: Auth Service, Product Service, and Order Service, each responsible for different functionalities of the e-commerce application.

## Project Structure

```
ecommerce-api
├── auth-service
│   ├── controllers
│   ├── dao
│   ├── models
│   ├── routes
│   └── main.go
├── product-service
│   ├── controllers
│   ├── dao
│   ├── models
│   ├── routes
│   └── main.go
├── order-service
│   ├── controllers
│   ├── dao
│   ├── models
│   ├── routes
│   └── main.go
├── go.mod
├── go.sum
└── README.md
```

## Services Overview

### Auth Service
- **Purpose**: Handles user authentication, registration, and session management.
- **Key Files**:
  - `controllers/auth_controller.go`: Contains methods for user authentication.
  - `dao/auth_dao.go`: Interacts with MongoDB for user data.
  - `models/auth_model.go`: Defines the User model.
  - `routes/auth_routes.go`: Sets up the authentication routes.
  - `main.go`: Entry point for the auth service.

### Product Service
- **Purpose**: Manages product listings, creation, and retrieval.
- **Key Files**:
  - `controllers/product_controller.go`: Contains methods for product management.
  - `dao/product_dao.go`: Interacts with MongoDB for product data.
  - `models/product_model.go`: Defines the Product model.
  - `routes/product_routes.go`: Sets up the product-related routes.
  - `main.go`: Entry point for the product service.

### Order Service
- **Purpose**: Manages customer orders and order history.
- **Key Files**:
  - `controllers/order_controller.go`: Contains methods for order management.
  - `dao/order_dao.go`: Interacts with MongoDB for order data.
  - `models/order_model.go`: Defines the Order model.
  - `routes/order_routes.go`: Sets up the order-related routes.
  - `main.go`: Entry point for the order service.

## Setup Instructions

1. **Clone the Repository**:
   ```
   git clone <repository-url>
   cd ecommerce-api
   ```

2. **Install Dependencies**:
   ```
   go mod tidy
   ```

3. **Run Services**:
   Each service can be run independently. Navigate to the service directory and execute:
   ```
   go run main.go
   ```

4. **Database Setup**:
   Ensure MongoDB is running and configured properly. Update the connection settings in the DAO files as needed.

## License
This project is licensed under the MIT License.