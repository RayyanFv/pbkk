// models/user.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the POS system
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"` // Hashed password
	Role      string         `json:"role"`                     // e.g., admin, cashier
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
