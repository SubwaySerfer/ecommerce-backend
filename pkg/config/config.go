package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port     string
    Database string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return &Config{
        Port:     getEnv("PORT", "8080"),
        Database: getEnv("DATABASE_URL", "your_default_database_url"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}