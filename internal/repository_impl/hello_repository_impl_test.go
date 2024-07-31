package repository_impl

import (
	"context"
	_ "github.com/mattn/go-sqlite3"

	"example.com/internal/infra/database/ent"
	"example.com/internal/infra/database/ent/enttest"
	"testing"
)

func setupTest(t *testing.T) (*repositoryContext, func()) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	ctx := context.Background()
	repo := &repositoryContext{
		client: client,
		ctx:    ctx,
	}
	return repo, func() {
		client.Close()
	}
}

func TestRepositoryContext_UpsertUser(t *testing.T) {
	repo, teardown := setupTest(t)
	defer teardown()

	if newUser, err := repo.UpsertUser(ent.User{Name: "John Doe", Age: 30}); err != nil || newUser.Name != "John Doe" || newUser.Age != 30 {
		t.Fatalf("failed upserting user: %v", err)
	} else {
		t.Logf("pass user upserted: %v", newUser)
	}
}

func TestRepositoryContext_UpsertUser_WithSameID(t *testing.T) {
	repo, teardown := setupTest(t)
	defer teardown()
	repo.UpsertUser(ent.User{ID: 1, Name: "John Doe", Age: 30})

	if newUser, err := repo.UpsertUser(ent.User{ID: 1, Name: "Anda Bool", Age: 20}); err != nil {
		t.Fatalf("failed upserting user: %v", err)
	} else {
		t.Logf("pass user upserted: %v", newUser)
	}
}

func TestRepositoryContext_GetUserById(t *testing.T) {
	repo, teardown := setupTest(t)
	defer teardown()
	repo.UpsertUser(ent.User{ID: 1, Name: "John Doe", Age: 30})

	if user, err := repo.GetUserById(1); err != nil || user.Name != "John Doe" || user.Age != 30 {
		t.Fatalf("failed getting user by id: %v", err)
	} else {
		t.Logf("pass user returned: %v", user)
	}
}

func TestRepositoryContext_GetUser_NotFound(t *testing.T) {
	repo, teardown := setupTest(t)
	defer teardown()
	repo.UpsertUser(ent.User{ID: 2, Name: "John Doe", Age: 30})

	if user, err := repo.GetUserById(1); err != nil {
		t.Logf("pass user not found: %v %v", err, user)
	} else {
		t.Fatalf("failed: %v %v", err, user)
	}
}
