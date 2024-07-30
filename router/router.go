package router

import (
	"example.com/internal/infra/database"
	"example.com/internal/repository_impl"
	"example.com/internal/service_impl"
	"example.com/router/handler"

	"github.com/gin-gonic/gin"
	"log"
)

func Start() {
	client := database.NewDatabaseClient()
	repository := repository_impl.NewRepository(client)
	service := service_impl.NewService(repository)
	apiHandler := handler.NewHandler(service)

	app := gin.Default()
	app.GET("/hello", apiHandler.Hello)

	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}
}
