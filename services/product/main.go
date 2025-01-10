package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"product-service/controllers"
	"product-service/services"
)

func main() {
	r := gin.Default()

	// Initialize MongoDB connection
	mongoService, err := services.NewMongoService()
	if err != nil {
		log.Fatal(err)
	}

	productController := controllers.NewProductController(mongoService)

	r.POST("/product", productController.CreateProduct)
	r.GET("/product/:id", productController.GetProduct)
	r.PUT("/product/:id", productController.UpdateProduct)
	r.DELETE("/product/:id", productController.DeleteProduct)

	r.Run(":8081")
}