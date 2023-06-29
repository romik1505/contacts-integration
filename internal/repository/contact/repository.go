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
	ListContacts(ctx context.Context, filter ListContactsFilter) ([]model.Contact, error)
	CreateContact(ctx context.Context, contact *model.Contact) error
	InsertContacts(ctx context.Context, contacts []model.Contact) (int64, error)
	UpdateContact(ctx context.Context, contact *model.Contact) error
	UpdateContactsByIDs(ctx context.Context, ids []uint64, c *model.Contact) (int64, error)
	UpdateContactsByAmoIDs(ctx context.Context, amoIDs []uint64, contact *model.Contact) error
	DeleteAccountContacts(ctx context.Context, accountID uint64) error
	DeleteContact(ctx context.Context, contact *model.Contact) error
	DeleteContacts(ctx context.Context, contacts []model.Contact) (int64, error)
	DeleteContactsByAmoIDs(ctx context.Context, amoIDs []uint64) (int64, error)
}

func NewRepository(s store.Store) *Repository {
	return &Repository{
		Store: s,
	}
}
