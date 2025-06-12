package handlers

import (
        "log"
        "net/http"
        "os"
        "strconv"
        "set-service/services"

        "github.com/gin-gonic/gin"
)

// ImportMinifigsHandler godoc
// @Summary Импорт минифигурок из Rebrickable
// @Description Загружает минифигурки частями и отправляет в data-service
// @Tags import
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string
// @Param page query int false "Страница начала"  default(1)
// @Param pages query int false "Количество страниц (0 - все)"  default(0)
// @Param page_size query int false "Размер страницы"  default(1000)
// @Router /api/import/minifigs [post]
func ImportMinifigsHandler(c *gin.Context) {
        apiKey := os.Getenv("REBRICKABLE_API_KEY")
        if apiKey == "" {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not set"})
                return
        }

        startPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
        pages, _ := strconv.Atoi(c.DefaultQuery("pages", "0"))
        pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "1000"))

        figs, err := services.FetchMinifigsChunk(apiKey, startPage, pageSize, pages)
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
