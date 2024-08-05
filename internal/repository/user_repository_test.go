package repository

import (
	"context"
	"example.com/internal/infra/database/ent/enttest"
	"example.com/internal/repository/model"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupTest(t *testing.T) (*userRepository, func()) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	ctx := context.Background()
	repo := &userRepository{
		client: client,
		ctx:    ctx,
	}
	return repo, func() {
		_ = client.Close()
	}
}

func TestUserRepository(t *testing.T) {
	t.Run("CreateUser", func(t *testing.T) {
		repo, teardown := setupTest(t)
		defer teardown()

		newUser, err := repo.CreateUser(model.UserInsertData{Name: "John Doe", Age: 30, Email: "a21"})
		if assert.NoError(t, err) {
			assert.Equal(t, newUser.Name, "John Doe")
			assert.Equal(t, newUser.Age, 30)
			assert.Equal(t, newUser.Email, "a21")
		}
	})
}
