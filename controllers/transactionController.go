// Handle creating a new transaction
package controllers

import (
	"fmt"
	"net/http"
	"pos-system/db"
	"pos-system/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handle creating a new transaction
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calculate the total amount and validate stock for each item in the transaction
	var totalAmount float64
	for _, item := range transaction.TransactionItems {
		var product models.Product
		if err := db.DB.First(&product, item.ProductID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Product with ID %d not found", item.ProductID)})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching product"})
			return
		}

		// Check if there is enough stock for the product
		if product.Stock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Not enough stock for product %s", product.Name)})
			return
		}

		// Calculate total price for the current item
		item.Price = product.Price
		item.TotalPrice = float64(item.Quantity) * product.Price
		totalAmount += item.TotalPrice
	}

	// Set the total amount for the transaction
	transaction.TotalAmount = totalAmount

	// Save the transaction
	if result := db.DB.Create(&transaction); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Save the transaction items
	for _, item := range transaction.TransactionItems {
		item.TransactionID = transaction.ID
		if err := db.DB.Create(&item).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction items"})
			return
		}
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction created successfully",
		"transaction": gin.H{
			"id":           transaction.ID,
			"total_amount": transaction.TotalAmount,
			"items":        transaction.TransactionItems,
		},
	})
}

// GetProductByID handles fetching a single product by ID
func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // Get the 'id' param from the URL
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Fetch the product from the database by its ID
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
