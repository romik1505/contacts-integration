package mapper

import (
	"fmt"
	"strconv"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/client/unisender"
	"week3_docker/internal/model"
	"week3_docker/internal/schemas"
	contact "week3_docker/pkg/api/contact_service"
)

func ConvertAmoContactsWithIDs(inp []schemas.Contact, accountID uint64, types string) ([]model.Contact, []uint64) {
	res := make([]model.Contact, 0, len(inp))
	ids := make([]uint64, 0)

	for _, contact := range inp {
		ids = append(ids, contact.ID)
		for _, fields := range contact.CustomFields {
			if fields.Code == "EMAIL" {
				for _, val := range fields.Values {
					cont := model.Contact{
						AccountID: accountID,
						AmoID:     contact.ID,
						Name:      contact.Name,
						Email:     val.Value,
						Type:      types,
						Sync:      false,
					}
					if cont.Valid() {
						res = append(res, cont)
					}
				}
			}
		}
	}
	return res, ids
}

func AmoContactsIDs(inp []schemas.Contact) []uint64 {
	ids := make([]uint64, len(inp))
	for i, c := range inp {
		ids[i] = c.ID
	}
	return ids
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
		AmoId:     c.AmoID,
		AccountId: c.AccountID,
		Name:      c.Name,
		Email:     c.Email,
		Type:      c.Type,
		Sync:      c.Sync,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ConvertAmoContactsToModel(cs []amo.Contact, types string) []model.Contact {
	res := make([]model.Contact, 0, 10)
	for _, v := range cs {
		for _, fields := range v.CustomFieldsValues {
			if fields.FieldCode == "EMAIL" {
				for _, field := range fields.Values {
					contact := model.Contact{
						AmoID:     uint64(v.ID),
						AccountID: uint64(v.AccountID),
						Name:      v.Name,
						Email:     field.Value,
						Type:      types,
						Sync:      false,
					}
					if contact.Valid() {
						res = append(res, contact)
					}
				}
			}
		}
	}
	return res
}

func BoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func ConvertContactsToUnisenderRequestParams(contacts []model.Contact, apiKey string, listID uint64) (unisender.ImportContactsRequest, error) {
	req := unisender.ImportContactsRequest{
		Format:     "json",
		ApiKey:     apiKey,
		FieldNames: []string{"email", "Name", "email_list_ids", "delete"},
	}
	req.Data = make([][]string, 0, len(contacts))
	for _, v := range contacts {
		if v.ReasonOutSync != "" {
			continue
		}
		req.Data = append(req.Data, []string{v.Email, v.Name, strconv.Itoa(int(listID)), BoolToString(v.Type == "delete")})
	}

	if len(req.Data) == 0 {
		return unisender.ImportContactsRequest{}, fmt.Errorf("convert to empty data")
	}

	return req, nil
}
