package handler

import (
	"example.com/internal/proto/payload"
	"example.com/internal/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

type UserHandler interface {
	GetUser(context echo.Context) error
	Register(context echo.Context) error
}

func (h *userHandler) GetUser(e echo.Context) error {
	idString := e.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("failed to convert id to int: %v", err)
		return e.JSON(http.StatusBadRequest, "Invalid id ")
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		log.Printf("failed to get user by id: %v", err)
		return e.JSON(http.StatusInternalServerError, "Failed to get user!")
	}

	return e.JSON(http.StatusOK, user)
}

func (h *userHandler) Register(e echo.Context) error {
	u := new(payload.UserRegisterPayload)
	err := e.Bind(u)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, "Register")
}
