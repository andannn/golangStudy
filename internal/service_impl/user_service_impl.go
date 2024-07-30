package service_impl

import (
	"example.com/proto"
)

func (s *serviceContext) GetUserById(id int) (*proto.UserResponse, error) {
	userEntity, err := s.repository.GetUserById(id)

	return proto.ToResponse(userEntity), err
}
