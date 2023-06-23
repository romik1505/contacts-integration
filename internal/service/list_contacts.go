package service

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"week3_docker/internal/mapper"
	"week3_docker/internal/model"
	contact "week3_docker/pkg/api/contact_service"
)

func (cs ContactService) ListContacts(ctx context.Context, request *contact.ListContactsRequest) (*contact.ListContactsResponse, error) {
	account, err := cs.ar.GetAccount(ctx, request.GetId())
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return new(contact.ListContactsResponse), nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp, err := cs.amoClient.ListContacts(ctx, account, model.ContactQueryParams{
		Page:  int(request.GetPage()),
		Limit: int(request.GetLimit()),
	})

	if err != nil {
		return nil, err
	}

	return &contact.ListContactsResponse{
		Items: mapper.ConvertContacts(resp.Embedded.Contacts),
	}, nil
}
