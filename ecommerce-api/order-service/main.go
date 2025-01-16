package main

import (
    "github.com/gin-gonic/gin"
    "ecommerce-api/order-service/routes"
)

func main() {
    r := gin.Default()
    routes.SetOrderRoutes(r)
    r.Run(":8080") // Run on port 8080
}