package repository_impl

import (
	"context"
	"example.com/domain/repository"
	"example.com/internal/infra/database/ent"
)

type repositoryContext struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepository(client *ent.Client) repository.Repository {
	return &repositoryContext{client: client}
}
