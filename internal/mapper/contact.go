package mapper

import (
	"week3_docker/internal/model"
	contact "week3_docker/pkg/api/contact_service"
)

func ConvertContacts(inp []model.Contact) []*contact.Contact {
	res := make([]*contact.Contact, 0, len(inp))
	for _, v := range inp {
		customFields := make([]*contact.CustomFieldsValue, len(v.CustomFieldsValues))
		for j, field := range v.CustomFieldsValues {
			customFieldsValues := make([]*contact.CustomFieldsValue_Values, len(field.Values))
			for k, val := range field.Values {
				customFieldsValues[k] = &contact.CustomFieldsValue_Values{
					Value:    val.Value,
					EnumId:   uint64(val.EnumID),
					EnumCode: val.EnumCode,
				}
			}
			customFields[j] = &contact.CustomFieldsValue{
				FieldId:   uint64(field.FieldID),
				FieldName: field.FieldName,
				FieldCode: field.FieldCode,
				FieldType: field.FieldType,
				Values:    customFieldsValues,
			}
		}
		res = append(res, &contact.Contact{
			Id:                uint64(v.ID),
			Name:              v.Name,
			FirstName:         v.FirstName,
			LastName:          v.LastName,
			ResponsibleUserId: uint64(v.ResponsibleUserID),
			CreatedAt:         uint64(v.CreatedAt),
			CreatedBy:         uint64(v.CreatedBy),
			UpdatedAt:         uint64(v.UpdatedAt),
			UpdatedBy:         uint64(v.UpdatedBy),
			IsDeleted:         v.IsDeleted,
			IsUnsorted:        v.IsUnsorted,
			CustomFieldValues: customFields,
			AccountId:         uint64(v.AccountID),
		})
	}
	return res
}
