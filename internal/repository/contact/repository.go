package contact

import (
	"context"
	"week3_docker/internal/model"
	"week3_docker/internal/store"
)

type Repository struct {
	Store store.Store
}

type IRepository interface {
	ListContacts(ctx context.Context, filter model.ListContactsFilter) ([]model.Contact, error)
	CreateContact(ctx context.Context, contact *model.Contact) error
	InsertContacts(ctx context.Context, contacts []*model.Contact) (int64, error)
	UpdateContact(ctx context.Context, contact *model.Contact) error
	DeleteAccountContacts(ctx context.Context, accountID uint64) error
}

func NewRepository(s store.Store) *Repository {
	return &Repository{
		Store: s,
	}
}
