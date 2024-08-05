package repository

import (
	"context"
	"example.com/internal/infra/database/ent"
	"example.com/internal/infra/database/ent/user"
	"example.com/internal/repository/model"
	"fmt"
	"log"
)

type userRepository struct {
	client *ent.Client
	ctx    context.Context
}

func NewUserRepository(client *ent.Client, ctx context.Context) UserRepository {
	return &userRepository{client: client, ctx: ctx}
}

type UserRepository interface {
	GetUserById(id int) (*ent.User, error)
	DeleteUserById(id int) error
	CreateUser(user model.UserInsertData) (*ent.User, error)
	GetAllUsers() ([]*ent.User, error)
}

func (r *userRepository) GetAllUsers() ([]*ent.User, error) {
	users, err := r.client.User.
		Query().
		All(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying users: %w", err)
	}
	log.Println("users returned: ", users)
	return users, nil
}

func (r *userRepository) GetUserById(id int) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.ID(id)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func (r *userRepository) DeleteUserById(id int) error {
	err := r.client.User.
		DeleteOneID(id).
		Exec(r.ctx)
	if err != nil {
		return fmt.Errorf("failed deleting user: %w", err)
	}
	log.Println("user deleted: ", id)
	return nil
}

func (r *userRepository) CreateUser(user model.UserInsertData) (*ent.User, error) {
	newUser, err := r.client.User.
		Create().
		SetName(user.Name).
		SetAge(user.Age).
		SetEmail(user.Email).
		Save(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed upserting user: %w", err)
	}
	return newUser, nil
}
