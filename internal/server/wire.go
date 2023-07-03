package server

import (
	"github.com/google/wire"
	"week3_docker/internal/handler"
	"week3_docker/internal/service/contact"
)

func InitializeServerExample() Server {
	wire.Build(
		NewServer,
		wire.Bind(new(contact.IService), new(*contact.Service)),
		handler.NewHandler,
		contact.ServiceSet,
	)
	return Server{}
}
