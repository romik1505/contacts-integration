package amo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"log"
	"net/http"
	"week3_docker/internal/errors"
	"week3_docker/internal/model"
)

const (
	contactsURLMask = "https://%s/api/v4/contacts"
)

func (a Client) ListContacts(ctx context.Context, account model.Account, params model.ContactQueryParams) (model.ContactResponse, error) {
	v, _ := query.Values(params)
	if !account.AccessToken.Valid || account.AccessToken.String == "" {
		return model.ContactResponse{}, fmt.Errorf("access_token not set")
	}
	if !account.Subdomain.Valid || account.Subdomain.String == "" {
		return model.ContactResponse{}, fmt.Errorf("subdomain not set")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(contactsURLMask+"?"+v.Encode(), account.Subdomain.String), nil)
	if err != nil {
		return model.ContactResponse{}, err
	}

	req.Header.Add("Authorization", "Bearer "+account.AccessToken.String)

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return model.ContactResponse{}, err
	}
	log.Println(resp.StatusCode)

	if resp.StatusCode < 200 || resp.StatusCode > 204 {
		var amoErr errors.AmoCRMError
		if err := json.NewDecoder(resp.Body).Decode(&amoErr); err != nil {
			return model.ContactResponse{}, err
		}
		return model.ContactResponse{}, fmt.Errorf("error response: %v", amoErr)
	}

	// 204 - контактов нет
	if resp.StatusCode == 204 {
		return model.ContactResponse{}, nil
	}

	var contactResponse model.ContactResponse
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&contactResponse); err != nil {
		log.Printf("json decode: %v", err)
		return model.ContactResponse{}, err
	}
	return contactResponse, nil
}
