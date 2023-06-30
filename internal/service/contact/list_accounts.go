package contact

import (
	"context"
	"fmt"
	"week3_docker/internal/mapper"
	"week3_docker/internal/repository/account"
	contact "week3_docker/pkg/api/contact_service"
)

func (s Service) ListAccounts(ctx context.Context, req *contact.ListAccountsRequest) (*contact.ListAccountsResponse, error) {
	accounts, err := s.ar.ListAccounts(ctx, account.ListAccountFilter{
		Page:          int(req.GetPage()),
		Limit:         int(req.GetLimit()),
		AmoAuthorized: req.AmoAuth,
	})
	if err != nil {
		return nil, fmt.Errorf("ListAccounts: %v", err)
	}

	return &contact.ListAccountsResponse{
		Items: mapper.ConvertAccounts(accounts),
	}, nil
}
