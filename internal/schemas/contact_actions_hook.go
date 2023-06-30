package schemas

import "strings"

type ContactActionsHookRequest struct {
	ID       uint64   `json:"id"`
	Account  Account  `json:"account"`
	Contacts Contacts `json:"contacts"`
}

type Contacts struct {
	Add    []Contact `json:"add"`
	Update []Contact `json:"update"`
	Delete []Contact `json:"delete"`
}

type Account struct {
	ID        uint64 `json:"id"`
	Subdomain string `json:"subdomain"`
	Links     Links  `json:"links"`
}

type Links struct {
	Self string `json:"_self"`
}

type Contact struct {
	ID                uint64        `json:"id"`
	Name              string        `json:"name"`
	ResponsibleUserID uint64        `json:"responsible_user_id"`
	DateCreate        uint64        `json:"date_create"`
	LastModified      uint64        `json:"last_modified"`
	CreatedUserID     uint64        `json:"created_user_id"`
	ModifiedUserID    uint64        `json:"modified_user_id"`
	CompanyName       string        `json:"company_name"`
	LinkedCompanyID   uint64        `json:"linked_company_id"`
	AccountID         uint64        `json:"account_id"`
	CustomFields      []CustomField `json:"custom_fields"`
	CreatedAt         uint64        `json:"created_at"`
	UpdatedAt         uint64        `json:"updated_at"`
	Type              string        `json:"type"`
}

type CustomField struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Values []struct {
		Value string `json:"value"`
		Enum  uint64 `json:"enum"`
	} `json:"values"`
	Code string `json:"code"`
}

func NextKey(s string) func() string {
	str := s
	last := false
	return func() string {
		if last {
			return ""
		}
		open := strings.Index(str, "[")
		closeBr := strings.Index(str, "]")
		if open == -1 || closeBr == -1 {
			last = true
			return str
		}
		curKey := str[0:open]

		nextKey := str[open+1 : closeBr]
		str = nextKey + str[closeBr+1:]
		return curKey
	}
}
