package repository_impl

import (
	"example.com/internal/infra/database/ent"
	"example.com/internal/infra/database/ent/user"
	"fmt"
	"log"
)

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
