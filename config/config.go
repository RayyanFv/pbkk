package config

import (
	"fmt"
	"log"
	"os"
	"pos-system/db"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Init initializes the DB connection
func Init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open connection to the database
	var err2 error
	db.DB, err2 = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err2 != nil {
		log.Fatalf("failed to connect to the database: %v", err2)
	}

	fmt.Println("Database connected successfully.")
}
