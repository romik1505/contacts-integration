package contact

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
	"week3_docker/internal/mapper"
	contact "week3_docker/pkg/api/contact_service"
)

func (s Service) GetAccount(ctx context.Context, req *contact.GetAccountRequest) (*contact.GetAccountResponse, error) {
	account, err := s.ar.GetAccount(ctx, req.GetId())

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, fmt.Errorf("GetAccount: %v", err)
	}
	return &contact.GetAccountResponse{
		Account: mapper.ConvertAccount(account),
	}, nil
}
