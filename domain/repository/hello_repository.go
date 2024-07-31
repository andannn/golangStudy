package repository

import (
	"example.com/internal/infra/database/ent"
)

type UserRepository interface {
	GetUserById(id int) (*ent.User, error)
	DeleteUserById(id int) error
	UpsertUser(user ent.User) (*ent.User, error)
	GetAllUsers() ([]*ent.User, error)
}
