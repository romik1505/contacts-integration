package amo

import (
	"context"
	"net/http"
	"week3_docker/internal/model"
)

type Client struct {
	httpClient http.Client
}

const (
	GrandTypeAccess  = "authorization_code"
	GrantTypeRefresh = "refresh_token"
)

type IAmoClient interface {
	ListContacts(ctx context.Context, account model.Account, params model.ContactQueryParams) (model.ContactResponse, error)
	AccessToken(ctx context.Context, account model.Account, request model.AuthRequest) (model.AuthTokenPair, error)
}

func NewAmoClient() *Client {
	return &Client{
		httpClient: http.Client{},
	}
}
