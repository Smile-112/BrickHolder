package db

import (
	"data-service/models"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		host, port, user, pass,
	)
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
		&models.Minifig{},
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	return nil
}
