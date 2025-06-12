// @title           BrickHolder.set-service
// @version         1.0
// @description     Сервис для работы с Rebrickable.
// @host            localhost:8080
// @BasePath        /
// @schemes 		http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"log"
	"os"
	_ "set-service/docs"
	"set-service/handlers"
	"set-service/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Загружаем .env файл (если он есть)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	apiKey := os.Getenv("REBRICKABLE_API_KEY")
	if apiKey == "" {
		log.Fatal("REBRICKABLE_API_KEY is not set in environment")
	}

	client := services.NewRebrickableClient(apiKey)
	setsHandler := handlers.NewSetsHandler(client)

	r := gin.Default()
	r.Use(cors.Default())
	api := r.Group("/api")
	{
		api.GET("/sets", setsHandler.GetSets)
		api.POST("/import/series", handlers.ImportSeriesHandler)
		api.POST("/import/sets", handlers.ImportSetsHandler)
		api.POST("/import/minifigs", handlers.ImportMinifigsHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}
