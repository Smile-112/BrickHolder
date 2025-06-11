// @title           BrickHolder.data-service
// @version         1.0
// @description     Сервис для работы с базой данных
// @host            localhost:8081
// @BasePath        /
// @schemes 		http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"data-service/db"
	"data-service/handlers"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	_ "data-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	_ = godotenv.Load()
	if err := db.Init(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	r := gin.Default()
	// Разрешаем все запросы с любого источника (для разработки)
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.POST("/lego/series", handlers.CreateSeriesHandler)
		api.GET("/lego/series", handlers.GetAllSeriesHandler)
		api.POST("/lego/sets", handlers.CreateSetHandler)
		api.GET("/lego/sets", handlers.GetAllSetHandler)
		api.GET("/users", handlers.GetAllUsersHandler)
		api.POST("/users", handlers.GetAllUsersHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	r.Run(":" + port)
}
