package service

import (
	"example.com/proto"
)

type userService interface {
	GetUserById(id int) (*proto.UserResponse, error)
}
