package unisender

import (
	"context"
	"net/http"
)

type Client struct {
	httpClient http.Client
}

const (
	UnisenderURL = "https://api.unisender.com"
)

type IUnisenderClient interface {
	GetLists(ctx context.Context, apiKey string) (GetListsResponse, error)
	CreateList(ctx context.Context, apiKey string, listTitle string) (uint64, error)
	DeleteList(ctx context.Context, apiKey string, listID uint64) error
	ImportContacts(ctx context.Context, req ImportContactsRequest) (ImportContactsResponse, error)
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{},
	}
}
