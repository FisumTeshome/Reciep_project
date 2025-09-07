package models

import (
	"github.com/google/uuid"
)

// RecipeIngredient represents the relationship between recipes and ingredients
type RecipeIngredient struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	RecipeID     uuid.UUID `json:"recipe_id" gorm:"type:uuid;not null"`
	IngredientID uuid.UUID `json:"ingredient_id" gorm:"type:uuid;not null"`
	Quantity     float64   `json:"quantity" gorm:"not null"`
	Unit         string    `json:"unit" gorm:"not null"`
	Notes        string    `json:"notes"`

	// Relationships
	Recipe     Recipe     `json:"recipe,omitempty" gorm:"foreignKey:RecipeID"`
	Ingredient Ingredient `json:"ingredient,omitempty" gorm:"foreignKey:IngredientID"`
}

// TableName specifies the table name for RecipeIngredient
func (RecipeIngredient) TableName() string {
	return "recipe_ingredients"
}
