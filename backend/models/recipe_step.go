package models

import (
	"time"

	"github.com/google/uuid"
)

// RecipeStep represents a step in a recipe
type RecipeStep struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	RecipeID     uuid.UUID `json:"recipe_id" gorm:"type:uuid;not null"`
	StepNumber   int       `json:"step_number" gorm:"not null"`
	Description  string    `json:"description" gorm:"not null"`
	Image        string    `json:"image"`
	TimeRequired int       `json:"time_required"` // in minutes
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Recipe Recipe `json:"recipe,omitempty" gorm:"foreignKey:RecipeID"`
}

// TableName specifies the table name for RecipeStep
func (RecipeStep) TableName() string {
	return "recipe_steps"
}
