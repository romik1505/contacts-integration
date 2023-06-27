package contact

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

func (cs Service) ListContacts(ctx context.Context, request *contact.ListContactsRequest) (*contact.ListContactsResponse, error) {
	contacts, err := cs.cr.ListContacts(ctx, model.ListContactsFilter{
		AccountID: request.GetId(),
		Page:      int(request.GetPage()),
		Limit:     int(request.GetLimit()),
		Type:      request.GetType(),
		Sync:      request.Sync,
	})

	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return new(contact.ListContactsResponse), nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &contact.ListContactsResponse{
		Items: mapper.ConvertContacts(contacts),
	}, nil
}
