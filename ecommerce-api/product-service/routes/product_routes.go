package routes

import (
    "github.com/gin-gonic/gin"
    "ecommerce-api/product-service/controllers"
)

func SetProductRoutes(router *gin.Engine) {
    productController := controllers.ProductController{}

    router.POST("/products", productController.CreateProduct)
    router.GET("/products/:id", productController.GetProduct)
    router.DELETE("/products/:id", productController.DeleteProduct)
}