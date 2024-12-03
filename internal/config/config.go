package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configurations
type Config struct {
	ServerPort string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("DB_HOST not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB_PORT not set")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB_USER not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("DB_PASSWORD not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME not set")
	}

	return &Config{
		ServerPort: serverPort,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbName:     dbName,
	}
}
