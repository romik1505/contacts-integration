package mapper

import (
	"week3_docker/internal/model"
	contact "week3_docker/pkg/api/contact_service"
)

func ConvertAccounts(inp []model.Account) []*contact.Account {
	res := make([]*contact.Account, len(inp))
	for i, v := range inp {
		res[i] = ConvertAccount(v)
	}
	return res
}

func ConvertAccount(inp model.Account) *contact.Account {
	return &contact.Account{
		Id:                 inp.ID,
		Subdomain:          inp.Subdomain,
		AmoAuth:            inp.AccessToken.String != "",
		UnisenderConnected: inp.UnisenderKey != "",
		CreatedAt:          inp.CreatedAt,
		UpdatedAt:          inp.UpdatedAt,
	}
}
