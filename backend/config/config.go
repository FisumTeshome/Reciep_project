package config

import (
    "log"
    "os"
	
    "github.com/joho/godotenv"
)

// Config holds the configuration values
type Config struct {
    Port     string
    Database string
}

// LoadConfig loads configuration from environment variables or .env file
func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }

    config := &Config{
        Port:     getEnv("PORT", "8080"),
        Database: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname"),
    }

    return config
}

// getEnv gets the environment variable or returns a default value
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}