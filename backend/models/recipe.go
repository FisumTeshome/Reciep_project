package models

import (
	"time"

	"github.com/google/uuid"
)

// Recipe represents a food recipe
type Recipe struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID        uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CategoryID    uuid.UUID `json:"category_id" gorm:"type:uuid;not null"`
	Title         string    `json:"title" gorm:"not null"`
	Description   string    `json:"description"`
	PrepTime      int       `json:"prep_time" gorm:"not null"` // in minutes
	CookTime      int       `json:"cook_time" gorm:"not null"` // in minutes
	Servings      int       `json:"servings" gorm:"not null"`
	Difficulty    string    `json:"difficulty" gorm:"default:'medium'"` // easy, medium, hard
	FeaturedImage string    `json:"featured_image" gorm:"not null"`
	Images        []string  `json:"images" gorm:"type:text[]"`
	IsPublished   bool      `json:"is_published" gorm:"default:false"`
	IsPaid        bool      `json:"is_paid" gorm:"default:false"`
	Price         float64   `json:"price" gorm:"default:0"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User              User               `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Category          Category           `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Steps             []RecipeStep       `json:"steps,omitempty" gorm:"foreignKey:RecipeID"`
	RecipeIngredients []RecipeIngredient `json:"recipe_ingredients,omitempty" gorm:"foreignKey:RecipeID"`
	Likes             []Like             `json:"likes,omitempty" gorm:"foreignKey:RecipeID"`
	Bookmarks         []Bookmark         `json:"bookmarks,omitempty" gorm:"foreignKey:RecipeID"`
	Comments          []Comment          `json:"comments,omitempty" gorm:"foreignKey:RecipeID"`
	Ratings           []Rating           `json:"ratings,omitempty" gorm:"foreignKey:RecipeID"`
}

// TableName specifies the table name for Recipe
func (Recipe) TableName() string {
	return "recipes"
}
