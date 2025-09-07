package models

import (
	"time"

	"github.com/google/uuid"
)

// Comment represents a user's comment on a recipe
type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	RecipeID  uuid.UUID `json:"recipe_id" gorm:"type:uuid;not null"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User   User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Recipe Recipe `json:"recipe,omitempty" gorm:"foreignKey:RecipeID"`
}

// TableName specifies the table name for Comment
func (Comment) TableName() string {
	return "comments"
}
