package model

import "fmt"

type ContactActionsTask struct {
	Type            string    `json:"-"`
	AccountID       uint64    `json:"account_id,omitempty"`
	UnisenderKey    string    `json:"unisender_key,omitempty"`
	UnisenderListID uint64    `json:"unisender_list_id,omitempty"`
	Contacts        []Contact `json:"contacts,omitempty"`
	IDs             []uint64  `json:"ids,omitempty"`
	TryNumber       int       `json:"try_number,omitempty"`
}

func (c ContactActionsTask) Validate() error {
	if c.AccountID == 0 {
		return fmt.Errorf("account_id not set")
	}
	if c.UnisenderKey == "" {
		return fmt.Errorf("unisender_key not set")
	}
	if len(c.IDs) == 0 {
		return fmt.Errorf("ids empty")
	}
	if len(c.Contacts) == 0 && (c.Type == "add" || c.Type == "update") {
		return fmt.Errorf("contacts empty")
	}
	return nil
}
