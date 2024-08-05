package service

import (
	"example.com/internal/proto/response"
	"example.com/internal/repository"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

type UserService interface {
	GetUserById(id int) (*response.UserResponse, error)
}

func (s *userService) GetUserById(id int) (*response.UserResponse, error) {
	userEntity, err := s.repository.GetUserById(id)

	return response.ToUseResponse(userEntity), err
}
