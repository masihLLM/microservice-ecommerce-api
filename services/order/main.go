package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"order-service/controllers"
	"order-service/services"
)

func main() {
	r := gin.Default()

	// Initialize MongoDB connection
	mongoService, err := services.NewMongoService()
	if err != nil {
		log.Fatal(err)
	}

	orderController := controllers.NewOrderController(mongoService)

	r.POST("/order", orderController.CreateOrder)
	r.GET("/order/:id", orderController.GetOrder)
	r.PUT("/order/:id", orderController.UpdateOrder)
	r.DELETE("/order/:id", orderController.DeleteOrder)

	r.Run(":8082")
}