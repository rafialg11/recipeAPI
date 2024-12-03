package config

import (
	"log"
	"os"
)

type Config struct {
	ServerPort string
	DbURL      string
}

func LoadConfig() *Config {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not set")
	}

	return &Config{
		ServerPort: serverPort,
		DbURL:      dbURL,
	}
}
