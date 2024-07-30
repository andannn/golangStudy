package service_impl

import (
	"example.com/domain/repository"
	"example.com/domain/service"
)

type serviceContext struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) service.Service {
	return &serviceContext{repository: repository}
}
