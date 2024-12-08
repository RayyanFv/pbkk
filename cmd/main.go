package main

import (
	"pos-system/config"
	"pos-system/controllers"
	"pos-system/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB connection
	config.Init()

	// Create Gin router
	router := gin.Default()

	// Enable CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Allow frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,                                                // Allow credentials (if needed)
		MaxAge:           12 * time.Hour,                                      // Cache preflight response for 12 hours
	}))

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

		auth.POST("/users", controllers.Register)
		auth.GET("/users", controllers.GetAllUsers)
		auth.PUT("/users/:id", controllers.UpdateUser)
		auth.DELETE("/users/:id", controllers.DeleteUser)
		// Product Routes
		auth.POST("/products", controllers.CreateProduct) // Create product
		auth.GET("/products", controllers.GetProducts)    // Get all products
		// auth.GET("/products/:id", controllers.GetTransactionByID) // Get product by ID
		auth.PUT("/products/:id", controllers.UpdateProduct)    // Update product
		auth.DELETE("/products/:id", controllers.DeleteProduct) // Delete Product by ID

		// Stock Routes
		auth.POST("/stocks", controllers.CreateStock) // Create stock for product

		// Transaction Routes
		auth.GET("/transactions", controllers.GetTransactions)
		auth.GET("/transactions/:id", controllers.GetTransactionByID)
		auth.POST("/transactions", controllers.CreateTransaction)
		// auth.GET("/transactions", controllers.GetTransactions)    // Get All Transactions
	}

	// Run server
	router.Run(":8080")
}
