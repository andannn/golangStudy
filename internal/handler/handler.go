package handler

import (
	"example.com/internal/service"
)

type Handler interface {
	UserHandler
}

type handlerContext struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handlerContext{service: service}
}
