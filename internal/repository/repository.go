package repository

import (
	"context"
	"example.com/internal/infra/database/ent"
)

type repositoryContext struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepository(client *ent.Client, ctx context.Context) Repository {
	return &repositoryContext{client: client, ctx: ctx}
}

type Repository interface {
	UserRepository
}
