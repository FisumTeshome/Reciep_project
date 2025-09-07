package models

import (
	"time"

	"github.com/google/uuid"
)

// Rating represents a user's rating on a recipe
type Rating struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	RecipeID  uuid.UUID `json:"recipe_id" gorm:"type:uuid;not null"`
	Rating    int       `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"` // 1-5 stars
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User   User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Recipe Recipe `json:"recipe,omitempty" gorm:"foreignKey:RecipeID"`
}

// TableName specifies the table name for Rating
func (Rating) TableName() string {
	return "ratings"
}
