// @title           BrickHolder.user-service
// @version         1.0
// @description     Сервис для пользовательских действий.
// @host            localhost:8080
// @BasePath        /
// @schemes 		http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Users

package main

import (
	_ "users-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	api := r.Group("/api/users/{user_token}")
	{
		api.GET("")
		api.POST("/import/series")
		api.POST("/import/sets")
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8085")
}
