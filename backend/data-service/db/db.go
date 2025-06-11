package db

import (
	"data-service/models"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	dsn := os.Getenv("DATABASE_URL") // или строишь строку из отдельных переменных
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Database connection established")

	// Авто создание таблиц
	err = DB.AutoMigrate(
		&models.Series{},
		&models.Set{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	return nil
}
