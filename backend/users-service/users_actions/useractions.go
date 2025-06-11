package usersactions

import "github.com/gin-gonic/gin"

// CreateSetListHandler godoc
// @Summary Создание списка наборов
// @Description создает новый список наборов для пользователя
// @Tags usersactions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/users/{user_token}/setlists [post]
func CreateSetListHandler(c *gin.Context) {}

// DelSetListHandler godoc
// @Summary Удаление списка наборов
// @Description Удаляет список наборов пользователя
// @Tags usersactions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/users/{user_token}/setlists [delete]
func DelSetListHandler(c *gin.Context) {}

// PatchSetListHandler godoc
// @Summary Обновляет пользовательский список наборов
// @Description Обновляет пользовательский список наборов
// @Tags usersactions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/users/{user_token}/setlists [patch]
func PatchSetListHandler(c *gin.Context) {}

// CreatePartListHandler godoc
// @Summary Создание списка деталей
// @Description создает новый список деталей для пользователя
// @Tags usersactions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/users/{user_token}/partlists [post]
func CreatePartListHandler(c *gin.Context) {}

// DelPartsListHandler godoc
// @Summary Удаление списка деталей
// @Description Удаляет список деталей пользователя
// @Tags usersactions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/users/{user_token}/partlists [delete]
func DelPartsListHandler(c *gin.Context) {}

// PatchPartsListHandler godoc
// @Summary Обновляет пользовательский список деталей
// @Description Обновляет пользовательский список деталей
// @Tags usersactions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/users/{user_token}/setlists [patch]
func PatchPartsListHandler(c *gin.Context) {}

// CreatePartListHandler godoc
// @Summary Создание списка минифигурок
// @Description создает новый список минифигурок для пользователя
// @Tags usersactions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/users/{user_token}/minifigslists [post]
func CreateMinifigsListHandler(c *gin.Context) {}

// DelMinifigsListHandler godoc
// @Summary Удаление списка минифигурок
// @Description Удаляет список минифигурок пользователя
// @Tags usersactions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/users/{user_token}/minifigslists [delete]
func DelMinifigsListHandler(c *gin.Context) {}

// PatchMinifigsListHandler godoc
// @Summary Обновляет пользовательский список минифигурок
// @Description Обновляет пользовательский список минифигурок
// @Tags usersactions
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Router /api/users/{user_token}/minifigslists [patch]
func PatchMinifigsListHandler(c *gin.Context) {}
