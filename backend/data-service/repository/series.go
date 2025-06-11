package repository

import (
	"data-service/db"
	"data-service/models"
	"log"
)

func AddSeries(series models.Series) error {
	if err := db.DB.Exec("DELETE FROM series").Error; err != nil {
		log.Println("Ошибка при удалении:", err)
	}
	return nil
}
