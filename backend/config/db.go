package config

import (
	"MilTrace/domain"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseURL() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	return dsn
}

func BootDB() (*gorm.DB, error) {
	url := GetDatabaseURL()
	db, err := gorm.Open(postgres.Open(url))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Enable UUID extension
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return nil, fmt.Errorf("failed to create uuid extension: %w", err)
	}

	// Migrate tables in correct order without FK constraints
	migrationOrder := []interface{}{&domain.Device{}, &domain.Location{}}
	for _, model := range migrationOrder {
		if err := db.AutoMigrate(model); err != nil {
			return nil, fmt.Errorf("failed to migrate %T: %w", model, err)
		}
	}

	log.Println("Database Initialized ✅")
	return db, nil
}
