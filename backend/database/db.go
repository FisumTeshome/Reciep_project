package database

import (
	"fmt"
	"log"
	"os"

	"github.com/FisumTeshome/Reciep_project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect establishes a connection to the PostgreSQL database
func Connect() {
	var err error

	// Get database connection details from environment variables
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "password"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "minabpro_recipes"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	// Create connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	// Configure GORM logger
	gormLogger := logger.Default.LogMode(logger.Info)

	// Connect to database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}

// AutoMigrate runs database migrations only when tables do not already exist.
// This avoids clashing with the SQL migrations that already created the schema.
func AutoMigrate() {
	modelsToEnsure := []interface{}{
		&models.User{},
		&models.Category{},
		&models.Ingredient{},
		&models.Recipe{},
		&models.RecipeStep{},
		&models.RecipeIngredient{},
		&models.Like{},
		&models.Bookmark{},
		&models.Comment{},
		&models.Rating{},
	}

	for _, model := range modelsToEnsure {
		if !DB.Migrator().HasTable(model) {
			if err := DB.Migrator().CreateTable(model); err != nil {
				log.Fatalf("Failed to create table for model %T: %v", model, err)
			}
			log.Printf("Created missing table for model %T", model)
		}
	}

	log.Println("Database schema check completed successfully")
}

// SeedDatabase seeds the database with initial data
func SeedDatabase() {
	// Seed categories
	categories := []models.Category{
		{Name: "Breakfast", Description: "Morning meals", Color: "#FF6B6B"},
		{Name: "Lunch", Description: "Midday meals", Color: "#4ECDC4"},
		{Name: "Dinner", Description: "Evening meals", Color: "#45B7D1"},
		{Name: "Dessert", Description: "Sweet treats", Color: "#96CEB4"},
		{Name: "Snacks", Description: "Quick bites", Color: "#FFEAA7"},
		{Name: "Beverages", Description: "Drinks and cocktails", Color: "#DDA0DD"},
		{Name: "Appetizers", Description: "Starters and small plates", Color: "#98D8C8"},
		{Name: "Soups", Description: "Warm and cold soups", Color: "#F7DC6F"},
		{Name: "Salads", Description: "Fresh and healthy", Color: "#82E0AA"},
		{Name: "Bread", Description: "Fresh baked goods", Color: "#F8C471"},
	}

	for _, category := range categories {
		var existingCategory models.Category
		if err := DB.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				DB.Create(&category)
				log.Printf("Created category: %s", category.Name)
			}
		}
	}

	log.Println("Database seeding completed")
}
