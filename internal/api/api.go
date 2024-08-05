package api

import (
	"context"
	"example.com/internal/handler"
	"example.com/internal/infra/database/ent"
	"example.com/internal/repository"
	"example.com/internal/service"
	"example.com/internal/util"
	"github.com/labstack/echo/v4"
)

type App struct {
	db *ent.Client
}

func NewApp(db *ent.Client) *App {
	return &App{
		db: db,
	}
}

func (app *App) Run() {
	r := repository.NewUserRepository(app.db, context.Background())
	s := service.NewUserService(r)
	userHandler := handler.NewUserHandler(s)

	e := echo.New()
	e.Validator = util.NewJSONValidator()
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/register", userHandler.Register)

	e.Logger.Fatal(e.Start(":1323"))
}
