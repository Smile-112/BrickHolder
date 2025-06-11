package repository

import (
	"data-service/db"
	"data-service/models"
	"log"
)

func AddSet(series models.Set) error {
	if err := db.DB.Exec("DELETE FROM set").Error; err != nil {
		log.Println("Ошибка при удалении:", err)
	}
	return nil
}
