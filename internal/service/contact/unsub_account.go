package contact

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	contact "week3_docker/pkg/api/contact_service"
)

func (s Service) UnsubAccount(ctx context.Context, req *contact.UnsubAccountRequest) (*emptypb.Empty, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id must be not 0")
	}
	err := s.ir.DeleteIntegrations(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	err = s.cr.DeleteAccountContacts(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	err = s.ar.DeleteAccount(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}
