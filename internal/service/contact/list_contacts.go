package contact

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"week3_docker/internal/mapper"
	contact_repo "week3_docker/internal/repository/contact"
	contact "week3_docker/pkg/api/contact_service"
)

func (cs Service) ListContacts(ctx context.Context, request *contact.ListContactsRequest) (*contact.ListContactsResponse, error) {
	contacts, err := cs.cr.ListContacts(ctx, contact_repo.ListContactsFilter{
		AccountID: request.GetId(),
		Page:      int(request.GetPage()),
		Limit:     int(request.GetLimit()),
		Type:      request.GetType(),
		Sync:      request.Sync,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &contact.ListContactsResponse{
		Items: mapper.ConvertContacts(contacts),
	}, nil
}
