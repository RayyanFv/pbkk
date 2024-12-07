package main

import (
	"pos-system/config"
	"pos-system/controllers"
	"pos-system/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB connection
	config.Init()

	// Create Gin router
	router := gin.Default()

	// Public route for login and register
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Root route that shows a simple message when visiting localhost
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the POS System!",
		})
	})

	// Authenticated routes group (using JWT authentication)
	auth := router.Group("/auth")
	auth.Use(utils.AuthMiddleware()) // Middleware for JWT token validation
	{
		// Sales Report Route (protected)
		auth.GET("/sales-report", controllers.GetSalesReport)

		// Product Routes
		auth.POST("/products", controllers.CreateProduct)       // Create product
		auth.GET("/products", controllers.GetProducts)          // Get all products
		auth.GET("/products/:id", controllers.GetProductByID)   // Get product by ID
		auth.PUT("/products/:id", controllers.UpdateProduct)    // Update product
		auth.DELETE("/products/:id", controllers.DeleteProduct) // Get Product by ID

		// Stock Routes
		auth.POST("/stocks", controllers.CreateStock) // Get Stock for Product

		// Transaction Routes
		auth.POST("/transactions", controllers.CreateTransaction) // Create Transaction
		// auth.GET("/transactions", controllers.GetTransactions)    // Get All Transactions
	}

	// Run server
	router.Run(":8080")
}
