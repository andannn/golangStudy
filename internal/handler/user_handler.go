package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type UserHandler interface {
	GetUser(context echo.Context) error
}

func (h *handlerContext) GetUser(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid id")
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		log.Printf("failed to get user by id: %v", err)
		return e.JSON(http.StatusInternalServerError, "Failed to get user!")
	}

	return e.JSON(200, user)
}
