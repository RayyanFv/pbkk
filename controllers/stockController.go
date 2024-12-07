// controllers/stockController.go
package controllers

import (
	"net/http"
	"pos-system/db"
	"pos-system/models"

	"github.com/gin-gonic/gin"
)

// CreateStock handles the creation of stock for a specific product
func CreateStock(c *gin.Context) {
	var stock models.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the product exists in the database
	var product models.Product
	if err := db.DB.First(&product, stock.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Insert new stock record into the database
	if err := db.DB.Create(&stock).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Stock created successfully",
		"stock":   stock,
	})
}
