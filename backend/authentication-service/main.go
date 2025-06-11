// @title           BrickHolder.authentication-service
// @version         1.0
// @description     Сервис авторизации
// @host            localhost:8082
// @BasePath        /
// @schemes 		http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	_ "autentification-service/docs"
	userActions "autentification-service/user-actions"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Для вывода в консоль

func main() {

	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/login", userActions.Login)
		api.POST("/register", userActions.Register)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8082")
}
