package mapper

import (
	"week3_docker/internal/model"
	contact "week3_docker/pkg/api/contact_service"
)

func ConvertAmoContacts(inp []*contact.AmoContact, accountID uint64, types string) []*model.Contact {
	res := make([]*model.Contact, 0, len(inp))
	for _, contact := range inp {
		for _, fields := range contact.GetCustomFields() {
			if fields.GetCode() == "EMAIL" {
				for _, val := range fields.GetValues() {
					email := val.GetValue()
					res = append(res, &model.Contact{
						AccountID: accountID,
						Name:      contact.GetName(),
						Email:     email,
						Type:      types,
						Sync:      false,
					})
				}
			}
		}
	}
	return res
}

func ConvertContacts(contacts []model.Contact) []*contact.Contact {
	res := make([]*contact.Contact, len(contacts))
	for i, v := range contacts {
		res[i] = ConvertContact(v)
	}
	return res
}

func ConvertContact(c model.Contact) *contact.Contact {
	return &contact.Contact{
		Id:        c.ID,
		AccountId: c.AccountID,
		Name:      c.Name,
		Email:     c.Email,
		Type:      c.Type,
		Sync:      c.Sync,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
