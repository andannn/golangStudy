package service

import (
	"example.com/internal/proto/response"
)

type userService interface {
	GetUserById(id int) (*response.UserResponse, error)
}

func (s *serviceContext) GetUserById(id int) (*response.UserResponse, error) {
	userEntity, err := s.repository.GetUserById(id)

	return response.ToUseResponse(userEntity), err
}
