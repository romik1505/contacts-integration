package unisender

import (
	"context"
	"week3_docker/internal/model"
)

type ImportContactsRequest struct {
	APIKey   string
	ListID   uint64
	Contacts []model.Contact
}

type ImportContactsResponse struct {
}

func (c Client) ImportContacts(ctx context.Context, req ImportContactsRequest) (ImportContactsResponse, error) {
	return ImportContactsResponse{}, nil
}
