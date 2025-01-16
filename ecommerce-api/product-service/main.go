package main

import (
    "github.com/gin-gonic/gin"
    "ecommerce-api/product-service/routes"
)

func main() {
    r := gin.Default()

    // Set up product routes
    routes.SetProductRoutes(r)

    // Start the server
    r.Run(":8080")
}