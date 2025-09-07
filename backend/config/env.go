package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file
func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }
}

// GetEnv gets the environment variable or returns a default value
func GetEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}