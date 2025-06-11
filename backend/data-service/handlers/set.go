package handlers

import (
	"data-service/db"
	"data-service/models"
	"net/http"

	//"data-service/db"
	"log"

	"github.com/gin-gonic/gin"
)

// CreateSetHandler godoc
// @Summary Добавить набор Lego
// @Description Добавляет набор в базу
// @Tags lego
// @Accept json
// @Produce json
// @Param set body models.Set true "Set data"
// @Success 201 "Запись успешно добавлена"
// @Failure 400
// @Failure 500
// @Router /api/lego/sets [post]
func CreateSetHandler(c *gin.Context) {
	var set models.Set
	if err := c.ShouldBindJSON(&set); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := db.DB.Create(&set).Error; err != nil {
		log.Printf("Failed to save set: %v", err)
		c.JSON(500, gin.H{"error": "Failed to save set"})
		return
	}

	c.JSON(201, set)
}

// GetAllSetHandler godoc
// @Summary      Get sets
// @Description  Возвращает список наборов
// @Router       /api/lego/sets [get]
// @Tags lego
func GetAllSetHandler(c *gin.Context) {
	var products []models.Set
	db.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}
