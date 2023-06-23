package contact

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/client/unisender"
	"week3_docker/internal/model"
	"week3_docker/internal/repository/account"
	contact_repository "week3_docker/internal/repository/contact"
	"week3_docker/internal/repository/integration"
	contact "week3_docker/pkg/api/contact_service"
)

type Service struct {
	amoClient amo.IAmoClient
	uniClient unisender.IUnisenderClient
	ar        account.IAccountRepository
	cr        contact_repository.IRepository
	ir        integration.IRepository

	contact.UnimplementedContactServiceServer
}

type IService interface {
	Login(ctx context.Context, account *model.Account) error
	AutoRefreshTokens()

	AuthIntegration(context.Context, *contact.AuthIntegrationRequest) (*emptypb.Empty, error)
	GetAccount(context.Context, *contact.GetAccountRequest) (*contact.GetAccountResponse, error)
	ListAccounts(context.Context, *contact.ListAccountsRequest) (*contact.ListAccountsResponse, error)
	ListContacts(context.Context, *contact.ListContactsRequest) (*contact.ListContactsResponse, error)
	ListAccountIntegrations(context.Context, *contact.ListAccountIntegrationsRequest) (*contact.ListAccountIntegrationsResponse, error)
	PrimaryContactsSync(context.Context, *contact.PrimaryContactSyncRequest) (*emptypb.Empty, error)
	ContactActionsHook(context.Context, *contact.ContactActionsHookRequest) (*emptypb.Empty, error)
	UnsubAccount(context.Context, *contact.UnsubAccountRequest) (*emptypb.Empty, error)
}

func NewService(
	amo amo.IAmoClient,
	uni unisender.IUnisenderClient,
	ar account.IAccountRepository,
	cr contact_repository.IRepository,
	ir integration.IRepository,
) *Service {
	return &Service{
		amoClient: amo,
		uniClient: uni,
		ar:        ar,
		cr:        cr,
		ir:        ir,
	}
}
