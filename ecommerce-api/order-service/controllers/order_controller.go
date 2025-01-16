package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "ecommerce-api/order-service/dao"
    "ecommerce-api/order-service/models"
)

type OrderController struct{}

func (oc *OrderController) CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := dao.CreateOrder(&order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
        return
    }
    c.JSON(http.StatusCreated, order)
}

func (oc *OrderController) GetOrder(c *gin.Context) {
    orderID := c.Param("id")
    order, err := dao.FindOrderByID(orderID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    c.JSON(http.StatusOK, order)
}

func (oc *OrderController) CancelOrder(c *gin.Context) {
    orderID := c.Param("id")
    if err := dao.CancelOrder(orderID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel order"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}