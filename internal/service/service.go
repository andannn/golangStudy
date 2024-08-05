package service

import "example.com/internal/repository"

type serviceContext struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &serviceContext{repository: repository}
}

type Service interface {
	userService
}
