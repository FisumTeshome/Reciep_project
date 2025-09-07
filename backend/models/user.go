package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username   string    `json:"username" gorm:"uniqueIndex;not null"`
	Email      string    `json:"email" gorm:"uniqueIndex;not null"`
	Password   string    `json:"-" gorm:"not null"` // "-" means don't include in JSON
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Avatar     string    `json:"avatar"`
	Bio        string    `json:"bio"`
	IsVerified bool      `json:"is_verified" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Recipes   []Recipe   `json:"recipes,omitempty" gorm:"foreignKey:UserID"`
	Likes     []Like     `json:"likes,omitempty" gorm:"foreignKey:UserID"`
	Bookmarks []Bookmark `json:"bookmarks,omitempty" gorm:"foreignKey:UserID"`
	Comments  []Comment  `json:"comments,omitempty" gorm:"foreignKey:UserID"`
	Ratings   []Rating   `json:"ratings,omitempty" gorm:"foreignKey:UserID"`
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "users"
}
