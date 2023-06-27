package contact

import (
	"context"
	"fmt"
	"week3_docker/internal/mapper"
	"week3_docker/internal/model"
	contact "week3_docker/pkg/api/contact_service"
)

func (s Service) ListAccountIntegrations(ctx context.Context, req *contact.ListAccountIntegrationsRequest) (*contact.ListAccountIntegrationsResponse, error) {
	integrations, err := s.ir.ListIntegration(ctx, model.ListIntegrationFilter{
		AccountID: int(req.GetId()),
		Page:      int(req.GetPage()),
		Limit:     int(req.GetLimit()),
	})
	if err != nil {
		return nil, fmt.Errorf("ListAccounts :%v", err)
	}

	return &contact.ListAccountIntegrationsResponse{
		Items: mapper.ConvertIntegrations(integrations),
	}, nil
}
