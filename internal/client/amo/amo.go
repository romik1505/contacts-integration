package amo

import (
	"context"
	"net/http"
)

type Client struct {
	httpClient http.Client
}

const (
	GrandTypeAccess  = "authorization_code"
	GrantTypeRefresh = "refresh_token"
)

type IAmoClient interface {
	ListContacts(ctx context.Context, request AccountRequest, params ContactQueryParams) (ContactResponse, error)
	AccessToken(ctx context.Context, subdomain string, req AuthRequest) (AuthTokenPair, error)
	Account(ctx context.Context, req AccountRequest) (AccountResponse, error)
	WebHookContactsSubscribe(ctx context.Context, request AccountRequest, accountID uint64) (WebhookSubscribeResponse, error)
}

func NewAmoClient() *Client {
	return &Client{
		httpClient: http.Client{},
	}
}
