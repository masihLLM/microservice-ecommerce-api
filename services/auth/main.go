package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"auth-service/controllers"
	"auth-service/services"
)

func main() {
	r := gin.Default()

	// Initialize MongoDB connection
	mongoService, err := services.NewMongoService()
	if err != nil {
		log.Fatal(err)
	}

	authController := controllers.NewAuthController(mongoService)

	r.POST("/auth/register", authController.Register)
	r.POST("/auth/login", authController.Login)

	r.Run(":8080")
}
