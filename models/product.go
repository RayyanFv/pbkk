package models

import (
	"gorm.io/gorm"
)

// Product model structure
type Product struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

// Create a new product in the database
func (p *Product) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

// Get all products from the database
func GetProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	err := db.Find(&products).Error
	return products, err
}

// Get a product by ID from the database
func GetProductByID(db *gorm.DB, id uint) (*Product, error) {
	var product Product
	err := db.First(&product, id).Error
	return &product, err
}

// Update an existing product in the database
func UpdateProduct(db *gorm.DB, p *Product) error {
	return db.Save(p).Error
}

// Delete a product by ID from the database
func DeleteProduct(db *gorm.DB, id uint) error {
	return db.Delete(&Product{}, id).Error
}
