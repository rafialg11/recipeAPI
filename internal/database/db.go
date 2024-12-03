package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm" // Or use github.com/gorm.io/gorm for newer versions of GORM
	_ "github.com/lib/pq"    // PostgreSQL driver
	"github.com/rafialg11/recipe-api/internal/config"
)

// ConnectDB initializes the database connection and returns a *gorm.DB object
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	// Build the database connection string
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)

	// Connect to the database
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return nil, err
	}

	// Auto migrate the database schema
	db.AutoMigrate()

	return db, nil
}
