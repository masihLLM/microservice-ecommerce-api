package routes

import (
	"github.com/gin-gonic/gin"
	"ecommerce-api/order-service/controllers"
)

func SetOrderRoutes(router *gin.Engine) {
	orderController := controllers.OrderController{}

	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders/:id", orderController.GetOrder)
	router.DELETE("/orders/:id", orderController.CancelOrder)
}