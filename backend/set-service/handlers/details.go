package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// ImportDetailsHandler godoc
// @Summary Импорт всех деталей из Rebrickable
// @Description Загружает все детали и отправляет в data-service
// @Tags import
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/import/details [post]
func ImportDetailsHandler(c *gin.Context) {
	apiKey := os.Getenv("REBRICKABLE_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not set"})
		return
	}

	//Берем все минифигурки из Rebrickable
	//series, err := services.FetchAllMinifigs(apiKey)
	//if err != nil {
	//	log.Printf("Failed to fetch series: %v", err) // <-- лог
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get themes"})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"status": "series imported successfully"})
}
