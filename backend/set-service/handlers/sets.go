package handlers

import (
	"log"
	"net/http"
	"os"
	_ "set-service/models"
	"set-service/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SetsHandler struct {
	Service *services.RebrickableClient
}

func NewSetsHandler(service *services.RebrickableClient) *SetsHandler {
	return &SetsHandler{Service: service}
}

// GetSets godoc
// @Summary      Получить список наборов LEGO
// @Description  Возвращает все наборы наборы с количеством, заданным ограничением
// @Tags         sets
// @Accept       json
// @Produce      json
// @Param        limit  query  int  false  "Количество наборов"
// @Success      200  {array}  models.Set
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/import/sets [post]
func (h *SetsHandler) GetSets(c *gin.Context) {
	limit := 10
	if l := c.Query("limit"); l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
		limit = val
	}

	sets, err := h.Service.GetLegoSets(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sets)
}

// ImportSetsHandler godoc
// @Summary Импорт всех наборов из Rebrickable
// @Description Загружает все наборы и отправляет в data-service
// @Tags import
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/import/sets [post]
func ImportSetsHandler(c *gin.Context) {
	apiKey := os.Getenv("REBRICKABLE_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not set"})
		return
	}

	//Берем все наборы из Rebrickable
	sets, err := services.FetchAllSets(apiKey)
	if err != nil {
		log.Printf("Failed to fetch sets: %v", err) // <-- лог
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sets"})
		return
	}

	//Передаем все наборы в микросервис БД
	for _, s := range sets {
		err := services.SendSetToDataService(s)
		if err != nil {
			log.Printf("Error sending set %s: %v", s.Name, err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "sets imported successfully"})
}
