package handlers

import (
	"data-service/db"
	"data-service/models"
	"net/http"

	//"data-service/db"
	"log"

	"github.com/gin-gonic/gin"
)

// CreateDetailsHandler godoc
// @Summary Создать деталь Lego
// @Description Добавляет деталь в базу
// @Tags lego
// @Accept json
// @Produce json
// @Success 201 "Запись успешно добавлена"
// @Failure 400
// @Failure 500
// @Router /api/lego/details [post]
func CreateDetailsHandler(c *gin.Context) {
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

// GetAllDetailsHandler godoc
// @Summary      Get details
// @Description  Возвращает список деталей
// @Router       /api/lego/details [get]
// @Tags lego
func GetAllDetailsHandler(c *gin.Context) {
	var products []models.Series
	db.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}
