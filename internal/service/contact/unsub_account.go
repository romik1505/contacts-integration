package contact

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	contact "week3_docker/pkg/api/contact_service"
)

func (s Service) UnsubAccount(ctx context.Context, req *contact.UnsubAccountRequest) (*emptypb.Empty, error) {
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
