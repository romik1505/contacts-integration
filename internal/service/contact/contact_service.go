package contact

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/client/unisender"
	"week3_docker/internal/model"
	"week3_docker/internal/queue"
	"week3_docker/internal/repository/account"
	contact_repository "week3_docker/internal/repository/contact"
	"week3_docker/internal/repository/integration"
	"week3_docker/internal/schemas"
	contact "week3_docker/pkg/api/contact_service"
)

type Service struct {
	amoClient amo.IAmoClient
	uniClient unisender.IUnisenderClient
	ar        account.IAccountRepository
	cr        contact_repository.IRepository
	ir        integration.IRepository

	queue *queue.Queue

	contact.UnimplementedContactServiceServer
}

type IService interface {
	Login(ctx context.Context, account *model.Account) error
	AutoRefreshTokens(ctx context.Context)
	InitSubscribeHook(ctx context.Context)

	AuthIntegration(context.Context, *contact.AuthIntegrationRequest) (*emptypb.Empty, error)
	GetAccount(context.Context, *contact.GetAccountRequest) (*contact.GetAccountResponse, error)
	ListAccounts(context.Context, *contact.ListAccountsRequest) (*contact.ListAccountsResponse, error)
	ListContacts(context.Context, *contact.ListContactsRequest) (*contact.ListContactsResponse, error)
	ListAccountIntegrations(context.Context, *contact.ListAccountIntegrationsRequest) (*contact.ListAccountIntegrationsResponse, error)

	// Pushed to queue
	PrimaryContactsSync(context.Context, schemas.ContactSyncRequest) error
	ContactActionsHook(context.Context, schemas.ContactActionsHookRequest) error

	// Do task
	DoPrimaryContactSync(ctx context.Context, request *contact.ContactSyncRequest) error
	DoAddContacts(ctx context.Context, task model.ContactActionsTask) error
	DoUpdateContacts(ctx context.Context, task model.ContactActionsTask) error
	DoDeleteContacts(ctx context.Context, task model.ContactActionsTask) error

	UnsubAccount(context.Context, *contact.UnsubAccountRequest) (*emptypb.Empty, error)
}

func NewService(
	amo amo.IAmoClient,
	uni unisender.IUnisenderClient,
	ar account.IAccountRepository,
	cr contact_repository.IRepository,
	ir integration.IRepository,
	q *queue.Queue,
) *Service {
	return &Service{
		amoClient: amo,
		uniClient: uni,
		ar:        ar,
		cr:        cr,
		ir:        ir,
		queue:     q,
	}
}

type Worker struct {
	ID      int
	Queue   *queue.Queue
	Service IService
}

func (s Service) Start() {
	workers := make([]Worker, 2)
	for i := range workers {
		workers[i] = Worker{
			ID:      i + 1,
			Queue:   s.queue,
			Service: s,
		}
	}

	for _, worker := range workers {
		go worker.DoTasks(context.Background())
	}

	go s.AutoRefreshTokens(context.Background())
	//go s.InitSubscribeHook(context.Background())
}
