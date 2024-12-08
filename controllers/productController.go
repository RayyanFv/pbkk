package controllers

import (
	"net/http"
	"pos-system/db"
	"pos-system/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create a product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Create method in the models layer
	if err := product.Create(db.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}

// Get all products
func GetProducts(c *gin.Context) {
	products, err := models.GetProducts(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Get a product by ID
func GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := models.GetProductByID(db.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Update a product
func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// First, get the product by ID to make sure it exists
	existingProduct, err := models.GetProductByID(db.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Now update the existing product data
	existingProduct.Name = product.Name
	existingProduct.Category = product.Category
	existingProduct.Price = product.Price
	existingProduct.Stock = product.Stock

	// Call the UpdateProduct function
	if err := models.UpdateProduct(db.DB, existingProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": existingProduct})
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Check if the product exists first
	_, err := models.GetProductByID(db.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Call the DeleteProduct function to remove it from the DB
	if err := models.DeleteProduct(db.DB, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
