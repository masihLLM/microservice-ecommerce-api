package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"ecommerce-api/product-service/dao"
	"ecommerce-api/product-service/models"
)

type ProductController struct{}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := dao.FindProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := dao.DeleteProduct(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}