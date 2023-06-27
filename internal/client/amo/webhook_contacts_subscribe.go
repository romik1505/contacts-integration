package amo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"week3_docker/internal/config"
	"week3_docker/internal/model"
)

const (
	subscribeHookUriMask = "https://%s.amocrm.ru/api/v4/webhooks"
	hookDestinationMask  = "%s/api/account/%d/contacts/hook"
)

func (c Client) WebHookContactsSubscribe(ctx context.Context, account model.Account, integration model.Integration) (model.WebhookSubscribeResponse, error) {
	if account.AccessToken == "" {
		return model.WebhookSubscribeResponse{}, fmt.Errorf("access for account id=%d not set", account.ID)
	}
	if account.Subdomain == "" {
		return model.WebhookSubscribeResponse{}, fmt.Errorf("subdomain for account id=%d not set", account.ID)
	}

	conString := fmt.Sprintf(subscribeHookUriMask, account.Subdomain)
	var buf bytes.Buffer

	uriDestinationHook := fmt.Sprintf(hookDestinationMask, config.Config.HostUrl, account.ID)
	err := json.NewEncoder(&buf).Encode(model.WebhookSubscribeRequest{
		Destinations: uriDestinationHook,
		Settings: []string{
			"restore_contact",
			"add_contact",
			"update_contact",
			"delete_contact",
		},
		Sort: 10,
	})
	if err != nil {
		return model.WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, conString, &buf)
	if err != nil {
		return model.WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+account.AccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return model.WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}

	respRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}

	var modelResp model.WebhookSubscribeResponse
	if err := json.Unmarshal(respRaw, &modelResp); err != nil {
		return model.WebhookSubscribeResponse{}, fmt.Errorf("WebHookContactsSubscribe: %v", err)
	}

	return modelResp, nil
}
