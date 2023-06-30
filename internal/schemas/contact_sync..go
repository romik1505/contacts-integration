package schemas

type ContactSyncRequest struct {
	AccountID    uint64 `schema:"account_id" json:"account_id,omitempty""`
	UnisenderKey string `schema:"unisender_key" json:"unisender_key,omitempty"`
}
