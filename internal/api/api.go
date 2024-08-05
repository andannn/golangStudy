package api

import (
	"context"
	"example.com/internal/handler"
	"example.com/internal/infra/database/ent"
	"example.com/internal/repository"
	"example.com/internal/service"
	"github.com/labstack/echo/v4"
)

type Server struct {
	db *ent.Client
}

func NewServer(db *ent.Client) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Run() {
	newRepository := repository.NewRepository(s.db, context.Background())
	newService := service.NewService(newRepository)
	apiHandler := handler.NewHandler(newService)

	e := echo.New()
	e.GET("/users/:id", apiHandler.GetUser)

	e.Logger.Fatal(e.Start(":1323"))
}
