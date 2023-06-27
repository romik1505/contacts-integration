package integration

import (
	"context"
	"week3_docker/internal/model"
	"week3_docker/internal/store"
)

type IRepository interface {
	CreateIntegration(ctx context.Context, integration *model.Integration) error
	UpdateIntegration(ctx context.Context, integration *model.Integration) error
	ListIntegration(ctx context.Context, filter model.ListIntegrationFilter) ([]model.Integration, error)
	GetIntegration(ctx context.Context, integration *model.Integration) (model.Integration, error)
	DeleteIntegrations(ctx context.Context, accountID uint64) error
}

type Repository struct {
	Store store.Store
}

func NewRepository(s store.Store) *Repository {
	return &Repository{
		Store: s,
	}
}
