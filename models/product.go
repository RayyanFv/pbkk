// models/product.go
package models

import (
	"gorm.io/gorm"
)

type Product struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

func (p *Product) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

func GetProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	err := db.Find(&products).Error
	return products, err
}

func GetProductByID(db *gorm.DB, id uint) (*Product, error) {
	var product Product
	err := db.First(&product, id).Error
	return &product, err
}

func UpdateProduct(db *gorm.DB, p *Product) error {
	return db.Save(p).Error
}

func DeleteProduct(db *gorm.DB, id uint) error {
	return db.Delete(&Product{}, id).Error
}
