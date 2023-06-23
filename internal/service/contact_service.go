package service

import (
	"context"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/model"
	"week3_docker/internal/repository/account"
	contact "week3_docker/pkg/api/contact_service"
)

type ContactService struct {
	amoClient amo.IAmoClient
	ar        account.IAccountRepository

	contact.UnimplementedContactServiceServer
}

type IContactService interface {
	Login(ctx context.Context, account *model.Account) error
	ListContacts(ctx context.Context, filter model.ContactQueryParams) ([]model.Contact, error)
	AutoRefreshTokens()
}

func NewContactService(amo amo.IAmoClient, ar account.IAccountRepository) *ContactService {
	return &ContactService{
		amoClient: amo,
		ar:        ar,
	}
}
