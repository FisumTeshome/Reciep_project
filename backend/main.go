package main

import (
	"log"
	"os"

	"github.com/FisumTeshome/Reciep_project/database"
	"github.com/FisumTeshome/Reciep_project/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize database connection
	database.Connect()

	// Run database migrations
	database.AutoMigrate()

	// Seed database with initial data
	database.SeedDatabase()

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting MinabPro Recipe API server on port %s...", port)

	// Initialize and start the server
	server := routes.InitializeRoutes()

	log.Printf("Server is running on http://localhost:%s", port)
	log.Printf("API Documentation available at http://localhost:%s/docs", port)

	// Start the server
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
