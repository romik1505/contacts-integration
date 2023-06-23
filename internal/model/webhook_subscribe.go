package model

type WebhookSubscribeRequest struct {
	Destinations string   `json:"destinations,omitempty"`
	Settings     []string `json:"settings,omitempty"`
	Sort         int      `json:"sort,omitempty"`
}

type WebhookSubscribeResponse struct {
	ID          int      `json:"id"`
	Destination string   `json:"destination"`
	CreatedAt   int      `json:"created_at"`
	UpdatedAt   int      `json:"updated_at"`
	AccountID   int      `json:"account_id"`
	CreatedBy   int      `json:"created_by"`
	Sort        int      `json:"sort"`
	Disabled    bool     `json:"disabled"`
	Settings    []string `json:"settings"`
}
