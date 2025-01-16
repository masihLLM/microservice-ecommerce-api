package main

import (
    "github.com/gin-gonic/gin"
    "ecommerce-api/auth-service/routes"
)

func main() {
    router := gin.Default()

    routes.SetAuthRoutes(router)

    router.Run(":8080")
}