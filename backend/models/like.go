package models

import (
	"time"

	"github.com/google/uuid"
)

// Like represents a user's like on a recipe
type Like struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	RecipeID  uuid.UUID `json:"recipe_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	User   User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Recipe Recipe `json:"recipe,omitempty" gorm:"foreignKey:RecipeID"`
}

// TableName specifies the table name for Like
func (Like) TableName() string {
	return "likes"
}
