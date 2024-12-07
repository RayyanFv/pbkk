// models/stock.go
package models

import (
	"gorm.io/gorm"
)

// Stock represents the stock of a product in the POS system
type Stock struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `json:"product_id"`
	Quantity  int            `json:"quantity"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
