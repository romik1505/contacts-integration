package amo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"week3_docker/internal/config"
)

const (
	subscribeHookUriMask = "https://%s/api/v4/webhooks"
	hookDestinationMask  = "%s/api/account/%d/contacts/hook"
)

type WebhookSubscribeRequest struct {
	Destination string   `json:"destination,omitempty"`
	Settings    []string `json:"settings,omitempty"`
	Sort        int      `json:"sort,omitempty"`
}

type WebhookSubscribeResponse struct {
	ID          int            `json:"id"`
	Destination string         `json:"destination"`
	CreatedAt   int            `json:"created_at"`
	UpdatedAt   int            `json:"updated_at"`
	AccountID   int            `json:"account_id"`
	CreatedBy   int            `json:"created_by"`
	Sort        int            `json:"sort"`
	Disabled    bool           `json:"disabled"`
	Settings    map[string]int `json:"settings"`
}

type WebhookSubscribeError struct {
	ValidationErrors []struct {
		RequestID string `json:"request_id"`
		Errors    []struct {
			Code   string `json:"code"`
			Path   string `json:"path"`
			Detail string `json:"detail"`
		} `json:"errors"`
	} `json:"validation-errors"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (w WebhookSubscribeError) Error() string {
	return w.Title
}

func (c Client) WebHookContactsSubscribe(ctx context.Context, request AccountRequest, accountID uint64) (WebhookSubscribeResponse, error) {
	if request.AccessToken == "" {
		return WebhookSubscribeResponse{}, fmt.Errorf("access for account id=%d not set", accountID)
	}
	if request.Subdomain == "" {
		return WebhookSubscribeResponse{}, fmt.Errorf("subdomain for account id=%d not set", accountID)
	}

	conString := fmt.Sprintf(subscribeHookUriMask, request.Subdomain)
	var buf bytes.Buffer

	uriDestinationHook := fmt.Sprintf(hookDestinationMask, config.Config.HostUrl, accountID)
	err := json.NewEncoder(&buf).Encode(WebhookSubscribeRequest{
		Destination: uriDestinationHook,
		Settings: []string{
			"restore_contact",
			"add_contact",
			"update_contact",
			"delete_contact",
		},
		Sort: 10,
	})
	if err != nil {
		return WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, conString, &buf)
	if err != nil {
		return WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+request.AccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}

	respRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var err WebhookSubscribeError
		if errUnmarshal := json.Unmarshal(respRaw, &err); errUnmarshal != nil {
			return WebhookSubscribeResponse{}, errUnmarshal
		}
		return WebhookSubscribeResponse{}, &err
	}

	var modelResp WebhookSubscribeResponse
	if err := json.Unmarshal(respRaw, &modelResp); err != nil {
		return WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}

	return modelResp, nil
}
