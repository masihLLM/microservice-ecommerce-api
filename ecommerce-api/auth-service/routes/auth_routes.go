package routes

import (
    "github.com/gin-gonic/gin"
    "ecommerce-api/auth-service/controllers"
)

func SetAuthRoutes(router *gin.Engine) {
    authController := controllers.AuthController{}

    router.POST("/auth/register", authController.CreateUser)
    router.POST("/auth/login", authController.Login)
    router.POST("/auth/logout", authController.Logout)
}