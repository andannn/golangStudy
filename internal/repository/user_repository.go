package repository

import (
	"example.com/internal/infra/database/ent"
	"example.com/internal/infra/database/ent/user"
	"fmt"
	"log"
)

type UserRepository interface {
	GetUserById(id int) (*ent.User, error)
	DeleteUserById(id int) error
	UpsertUser(user ent.User) (*ent.User, error)
	GetAllUsers() ([]*ent.User, error)
}

func (r *repositoryContext) GetAllUsers() ([]*ent.User, error) {
	users, err := r.client.User.
		Query().
		All(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying users: %w", err)
	}
	log.Println("users returned: ", users)
	return users, nil
}

func (r *repositoryContext) GetUserById(id int) (*ent.User, error) {
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

func (r *repositoryContext) DeleteUserById(id int) error {
	err := r.client.User.
		DeleteOneID(id).
		Exec(r.ctx)
	if err != nil {
		return fmt.Errorf("failed deleting user: %w", err)
	}
	log.Println("user deleted: ", id)
	return nil
}

func (r *repositoryContext) UpsertUser(user ent.User) (*ent.User, error) {
	newUser, err := r.client.User.
		Create().
		SetName(user.Name).
		SetAge(user.Age).
		Save(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed upserting user: %w", err)
	}
	log.Println("user upserted: ", user)
	return newUser, nil
}
