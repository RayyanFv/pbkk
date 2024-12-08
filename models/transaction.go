package models

import (
	"time"

	"gorm.io/gorm"
)

// TransactionItem represents an item in a transaction, linking product and quantity.
type TransactionItem struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	TransactionID uint    `json:"transaction_id"` // Foreign key to the Transaction table
	ProductID     uint    `json:"product_id"`
	Product       Product `json:"product" gorm:"foreignKey:ProductID"` // Join with Product model
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"` // Store price at the time of transaction
	TotalPrice    float64 `json:"total_price"`
}

// Transaction represents a transaction in the POS system.
// type Transaction struct {
// 	ID               uint              `json:"id" gorm:"primaryKey"`
// 	CustomerID       uint              `json:"customer_id"`
// 	TotalAmount      float64           `json:"total_amount"`
// 	TransactionItems []TransactionItem `json:"items" gorm:"foreignKey:TransactionID"` // One-to-many relationship with TransactionItem
// 	CreatedAt        time.Time         `json:"created_at"`
// 	UpdatedAt        time.Time         `json:"updated_at"`
// 	DeletedAt        gorm.DeletedAt    `json:"deleted_at,omitempty"`
// }

type Transaction struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CustomerID  uint           `json:"customer_id"`
	TotalAmount float64        `json:"total_amount"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (TransactionItem) TableName() string {
	return "transaction_item" // Sesuaikan dengan nama tabel yang ada di database
}
