package handlers

import (
	"log"
	"net/http"
	"os"
	"set-service/services"

	"github.com/gin-gonic/gin"
)

// ImportSeriesHandler godoc
// @Summary Импорт всех серий из Rebrickable
// @Description Загружает все серии и отправляет в data-service
// @Tags import
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/import/series [post]
func ImportSeriesHandler(c *gin.Context) {
	apiKey := os.Getenv("REBRICKABLE_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not set"})
		return
	}

	//Берем все серии из Rebrickable
	series, err := services.FetchAllSeries(apiKey)
	if err != nil {
		log.Printf("Failed to fetch series: %v", err) // <-- лог
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get themes"})
		return
	}

	//Передаем все серии в микросервис БД
	for _, s := range series {
		err := services.SendSeriesToDataService(s)
		if err != nil {
			log.Printf("Error sending series %s: %v", s.Name, err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "series imported successfully"})
}
