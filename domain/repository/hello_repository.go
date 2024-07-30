package repository

import (
	"example.com/internal/infra/database/ent"
)

type UserRepository interface {
	GetUserById(id int) (*ent.User, error)
}
