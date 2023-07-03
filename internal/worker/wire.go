package worker

import (
	"github.com/google/wire"
	"week3_docker/internal/service/contact"
)

func InitializeWorkerPoolExample() WorkerPool {
	wire.Build(
		NewWorkerPool,
		contact.ServiceSet,
	)
	return WorkerPool{}
}
