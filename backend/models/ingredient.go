package models

import (
	"time"

	"github.com/google/uuid"
)

// Ingredient represents an ingredient used in recipes
type Ingredient struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description"`
	Unit        string    `json:"unit"` // e.g., "grams", "cups", "pieces"
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	RecipeIngredients []RecipeIngredient `json:"recipe_ingredients,omitempty" gorm:"foreignKey:IngredientID"`
}

// TableName specifies the table name for Ingredient
func (Ingredient) TableName() string {
	return "ingredients"
}
