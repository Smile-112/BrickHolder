package handlers

import (
	"log"
	"net/http"
	"os"
	"set-service/services"

	"github.com/gin-gonic/gin"
)

// ImportMinifigsHandler godoc
// @Summary Импорт всех минифигурок из Rebrickable
// @Description Загружает все минифигурки и отправляет в data-service
// @Tags import
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/import/minifigs [post]
func ImportMinifigsHandler(c *gin.Context) {
	apiKey := os.Getenv("REBRICKABLE_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not set"})
		return
	}

	figs, err := services.FetchAllMinifigs(apiKey)
	if err != nil {
		log.Printf("Failed to fetch minifigs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get minifigs"})
		return
	}

	for _, f := range figs {
		if err := services.SendMinifigToDataService(f); err != nil {
			log.Printf("Error sending minifig %s: %v", f.Name, err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "minifigs imported successfully"})
}
