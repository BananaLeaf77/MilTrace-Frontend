package config

import (
	"MilTrace/domain"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseURL() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	return dsn
}

func BootDB() (*gorm.DB, error) {
	url := GetDatabaseURL()
	var err error

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.AutoMigrate(&domain.Device{}, &domain.User{})

	fmt.Println("DB initialized")
	return db, nil
}
