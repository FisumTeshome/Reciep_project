package models

import (
	"time"

	"github.com/google/uuid"
)

// Category represents a food category
type Category struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Color       string    `json:"color"` // For UI styling
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Recipes []Recipe `json:"recipes,omitempty" gorm:"foreignKey:CategoryID"`
}

// TableName specifies the table name for Category
func (Category) TableName() string {
	return "categories"
}
