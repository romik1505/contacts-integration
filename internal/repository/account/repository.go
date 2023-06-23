package account

import (
	"context"
	"week3_docker/internal/model"
	"week3_docker/internal/store"
)

type IAccountRepository interface {
	CreateAccount(ctx context.Context, account *model.Account) error
	UpdateAccount(ctx context.Context, account *model.Account) error
	DeleteAccount(ctx context.Context, id uint64) error
	ListAccounts(ctx context.Context, filter model.ListAccountFilter) ([]model.Account, error)
	GetAccount(ctx context.Context, id uint64) (model.Account, error)
}

type Repository struct {
	Store store.Store
}

func NewAccountRepository(store store.Store) *Repository {
	return &Repository{
		Store: store,
	}
}
