package contact

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"week3_docker/internal/model"
	contact "week3_docker/pkg/api/contact_service"
)

func (cs Service) AuthIntegration(ctx context.Context, request *contact.AuthIntegrationRequest) (*emptypb.Empty, error) {
	account := &model.Account{
		Subdomain: request.GetReferer(),
		AuthCode:  request.GetCode(),
		Integrations: []model.Integration{
			{
				OuterID: request.GetClientId(),
			},
		},
	}

	err := cs.ar.CreateAccount(ctx, account)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	go func() {
		// TODO: context canceled
		err := cs.Login(context.Background(), account)
		if err != nil {
			log.Printf("err auth: %v", err)
		}
	}()

	return new(emptypb.Empty), err
}
