package handlers

import (
	"data-service/db"
	"data-service/models"
	"net/http"

	//"data-service/db"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateMinifigsHandler godoc
// @Summary Создать минифигурку Lego
// @Description Добавляет минифигурку в базу
// @Tags lego
// @Accept json
// @Produce json
// @Success 201 "Запись успешно добавлена"
// @Failure 400
// @Failure 500
// @Router /api/lego/minifigs [post]
func CreateMinifigsHandler(c *gin.Context) {
	var fig models.Minifig
	if err := c.ShouldBindJSON(&fig); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var existing models.Minifig
	err := db.DB.First(&existing, "set_num = ?", fig.SetNum).Error
	if err == nil {
		c.JSON(http.StatusOK, existing)
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Printf("Failed to check existing minifig: %v", err)
		c.JSON(500, gin.H{"error": "Failed to save minifig"})
		return
	}

	if err := db.DB.Create(&fig).Error; err != nil {
		log.Printf("Failed to save minifig: %v", err)
		c.JSON(500, gin.H{"error": "Failed to save minifig"})
		return
	}

	c.JSON(201, fig)
}

// GetAllMinifigsHandler godoc
// @Summary      Get minifigs
// @Description  Возвращает список минифигурок
// @Router       /api/lego/minifigs [get]
// @Tags lego
func GetAllMinifigsHandler(c *gin.Context) {
	var figs []models.Minifig
	db.DB.Find(&figs)
	c.JSON(http.StatusOK, gin.H{
		"data": figs,
	})
}
