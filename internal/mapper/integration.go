package mapper

import (
	"week3_docker/internal/model"
	contact "week3_docker/pkg/api/contact_service"
)

func ConvertIntegrations(inp []model.Integration) []*contact.AccountIntegration {
	res := make([]*contact.AccountIntegration, len(inp))
	for i, v := range inp {
		res[i] = ConvertIntegration(v)
	}
	return res
}

func ConvertIntegration(inp model.Integration) *contact.AccountIntegration {
	return &contact.AccountIntegration{
		Id:        inp.ID,
		OuterId:   inp.OuterID,
		CreatedAt: inp.CreatedAt,
		UpdatedAt: inp.UpdatedAt,
	}
}
