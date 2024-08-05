package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type JSONValidator interface {
	Validate(i interface{}) error
}

type jsonValidator struct {
	validator *validator.Validate
}

func NewJSONValidator() JSONValidator {
	return &jsonValidator{validator: validator.New()}
}

func (jv *jsonValidator) Validate(i interface{}) error {
	if err := jv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
