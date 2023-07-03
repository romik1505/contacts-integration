package contact

import (
	"github.com/google/wire"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/client/unisender"
	"week3_docker/internal/queue"
	"week3_docker/internal/repository/account"
	contact_repository "week3_docker/internal/repository/contact"
	"week3_docker/internal/repository/integration"
	"week3_docker/internal/store"
)

var ServiceSet = wire.NewSet(
	NewService,
	amo.NewAmoClient,
	wire.Bind(new(amo.IAmoClient), new(*amo.Client)),

	unisender.NewClient,
	wire.Bind(new(unisender.IUnisenderClient), new(*unisender.Client)),

	account.NewAccountRepository,
	wire.Bind(new(account.IAccountRepository), new(*account.Repository)),

	contact_repository.NewRepository,
	wire.Bind(new(contact_repository.IRepository), new(*contact_repository.Repository)),

	integration.NewRepository,
	wire.Bind(new(integration.IRepository), new(*integration.Repository)),

	store.NewStore,
	queue.NewQueue,
)
