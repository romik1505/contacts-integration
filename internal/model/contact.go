package model

type Contact struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	FirstName          string      `json:"first_name"`
	LastName           string      `json:"last_name"`
	ResponsibleUserID  int         `json:"responsible_user_id"`
	GroupID            int         `json:"group_id"`
	CreatedBy          int         `json:"created_by"`
	UpdatedBy          int         `json:"updated_by"`
	CreatedAt          int         `json:"created_at"`
	UpdatedAt          int         `json:"updated_at"`
	ClosestTaskAt      interface{} `json:"closest_task_at"`
	IsDeleted          bool        `json:"is_deleted"`
	IsUnsorted         bool        `json:"is_unsorted"`
	CustomFieldsValues []struct {
		FieldID   int    `json:"field_id"`
		FieldName string `json:"field_name"`
		FieldCode string `json:"field_code"`
		FieldType string `json:"field_type"`
		Values    []struct {
			Value    string `json:"value"`
			EnumID   int    `json:"enum_id"`
			EnumCode string `json:"enum_code"`
		} `json:"values"`
	} `json:"custom_fields_values"`
	AccountID int `json:"account_id"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Embedded struct {
		Tags      []interface{} `json:"tags"`
		Companies []struct {
			ID    int `json:"id"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"companies"`
	} `json:"_embedded"`
}

type ContactResponse struct {
	Page  int `json:"_page"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Embedded struct {
		Contacts []Contact `json:"contacts"`
	} `json:"_embedded"`
}

type ContactQueryParams struct {
	With   string        `schema:"with" url:"with"`
	Page   int           `schema:"page" url:"page"`
	Limit  int           `schema:"limit" url:"limit"`
	Filter ContactFilter `schema:"filter" url:"filter"`
	Order  ContactOrder  `schema:"order" url:"order"`
}

type ContactFilter struct {
	ID                string `schema:"id" url:"id"`
	Name              string `schema:"name" url:"name"`
	CreatedBy         int64  `schema:"created_by" url:"created_by"`
	UpdatedBy         int64  `schema:"updated_by" url:"updated_by"`
	ResponsibleUserID int64  `schema:"responsible_user_id" url:"responsible_user_id"`
}

type ContactOrder struct {
	ID        string `schema:"id" url:"id"`
	UpdatedAt string `schema:"updated_at" url:"updated_at"`
}
