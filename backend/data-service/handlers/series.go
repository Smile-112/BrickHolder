package handlers

import (
	"data-service/db"
	"data-service/models"
	"net/http"

	//"data-service/db"
	"log"

	"github.com/gin-gonic/gin"
)

// CreateSeriesHandler godoc
// @Summary Создать серию Lego
// @Description Добавляет серию в базу
// @Tags lego
// @Accept json
// @Produce json
// @Param series body models.Series true "Series data"
// @Success 201 "Запись успешно добавлена"
// @Failure 400
// @Failure 500
// @Router /api/lego/series [post]
func CreateSeriesHandler(c *gin.Context) {
	var series models.Series
	if err := c.ShouldBindJSON(&series); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := db.DB.Create(&series).Error; err != nil {
		log.Printf("Failed to save series: %v", err)
		c.JSON(500, gin.H{"error": "Failed to save series"})
		return
	}

	c.JSON(201, series)
}

// GetAllSeriesHandler godoc
// @Summary      Get series
// @Description  Возвращает список серий
// @Router       /api/lego/series [get]
// @Tags lego
func GetAllSeriesHandler(c *gin.Context) {
	var series []models.Series
	db.DB.Find(&series)
	c.JSON(http.StatusOK, gin.H{
		"data": series,
	})
}
