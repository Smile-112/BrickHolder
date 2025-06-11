package handlers

import (
	"data-service/db"
	"data-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"Логин"`
	Password string `json:"Пароль"`
}
type TokenResponse struct {
	Token string `json:"token"`
}

// CreateUserHandler godoc
// @Summary Создать пользователя
// @Description Добавляет пользователя в базу при регистрации
// @Tags users
// @Accept json
// @Produce json
// @Success 201 "Запись успешно добавлена"
// @Param series body models.User true "Series data"
// @Failure 400
// @Failure 500
// @Router /api/users [post]
func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}

// GetAllDetailsHandler godoc
// @Summary      Получить список всех пользователей
// @Description  Возвращает список всех пользователей
// @Router       /api/users [get]
// @Tags users
func GetAllUsersHandler(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
