package amo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const accountURLMask = "https://%s/api/v4/account"

type AccountRequest struct {
	Subdomain   string
	AccessToken string
}

type AccountResponse struct {
	ID                      uint64 `json:"id"`
	Name                    string `json:"name"`
	Subdomain               string `json:"subdomain"`
	CreatedAt               int    `json:"created_at"`
	CreatedBy               int    `json:"created_by"`
	UpdatedAt               int    `json:"updated_at"`
	UpdatedBy               int    `json:"updated_by"`
	CurrentUserID           int    `json:"current_user_id"`
	Country                 string `json:"country"`
	Currency                string `json:"currency"`
	CurrencySymbol          string `json:"currency_symbol"`
	CustomersMode           string `json:"customers_mode"`
	IsUnsortedOn            bool   `json:"is_unsorted_on"`
	MobileFeatureVersion    int    `json:"mobile_feature_version"`
	IsLossReasonEnabled     bool   `json:"is_loss_reason_enabled"`
	IsHelpbotEnabled        bool   `json:"is_helpbot_enabled"`
	IsTechnicalAccount      bool   `json:"is_technical_account"`
	ContactNameDisplayOrder int    `json:"contact_name_display_order"`
	Links                   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func (c Client) Account(ctx context.Context, request AccountRequest) (AccountResponse, error) {
	if request.Subdomain == "" {
		return AccountResponse{}, fmt.Errorf("subdomain not set")
	}
	if request.AccessToken == "" {
		return AccountResponse{}, fmt.Errorf("account unauthorized")
	}

	url := fmt.Sprintf(accountURLMask, request.Subdomain)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return AccountResponse{}, err
	}

	req.Header.Add("Authorization", "Bearer "+request.AccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AccountResponse{}, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AccountResponse{}, err
	}

	var res AccountResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return AccountResponse{}, err
	}
	return res, nil
}
