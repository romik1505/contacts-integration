package amo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"log"
	"net/http"
	"week3_docker/internal/errors"
)

const (
	contactsURLMask = "https://%s/api/v4/contacts"
)

type ContactQueryParams struct {
	With   string        `schema:"with" url:"with"`
	Page   int           `schema:"page" url:"page"`
	Limit  int           `schema:"limit" url:"limit"`
	Filter ContactFilter `schema:"filter" url:"filter"`
	Order  ContactOrder  `schema:"order" url:"order"`
}

type ContactFilter struct {
	ID                string `schema:"id" url:"id"`
	Name              string `schema:"name" url:"name"`
	CreatedBy         int64  `schema:"created_by" url:"created_by"`
	UpdatedBy         int64  `schema:"updated_by" url:"updated_by"`
	ResponsibleUserID int64  `schema:"responsible_user_id" url:"responsible_user_id"`
}

type ContactOrder struct {
	ID        string `schema:"id" url:"id"`
	UpdatedAt string `schema:"updated_at" url:"updated_at"`
}

type Contact struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	FirstName          string      `json:"first_name"`
	LastName           string      `json:"last_name"`
	ResponsibleUserID  int         `json:"responsible_user_id"`
	GroupID            int         `json:"group_id"`
	CreatedBy          int         `json:"created_by"`
	UpdatedBy          int         `json:"updated_by"`
	CreatedAt          int         `json:"created_at"`
	UpdatedAt          int         `json:"updated_at"`
	ClosestTaskAt      interface{} `json:"closest_task_at"`
	IsDeleted          bool        `json:"is_deleted"`
	IsUnsorted         bool        `json:"is_unsorted"`
	CustomFieldsValues []struct {
		FieldID   int    `json:"field_id"`
		FieldName string `json:"field_name"`
		FieldCode string `json:"field_code"`
		FieldType string `json:"field_type"`
		Values    []struct {
			Value    string `json:"value"`
			EnumID   int    `json:"enum_id"`
			EnumCode string `json:"enum_code"`
		} `json:"values"`
	} `json:"custom_fields_values"`
	AccountID int `json:"account_id"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Embedded struct {
		Tags      []interface{} `json:"tags"`
		Companies []struct {
			ID    int `json:"id"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"companies"`
	} `json:"_embedded"`
}

type ContactResponse struct {
	Page  int `json:"_page"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Embedded struct {
		Contacts []Contact `json:"contacts"`
	} `json:"_embedded"`
}

// ListContacts access and subdomain required
func (a Client) ListContacts(ctx context.Context, request AccountRequest, params ContactQueryParams) (ContactResponse, error) {
	v, _ := query.Values(params)
	if request.AccessToken == "" {
		return ContactResponse{}, fmt.Errorf("access_token not set")
	}
	if request.Subdomain == "" {
		return ContactResponse{}, fmt.Errorf("subdomain not set")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(contactsURLMask+"?"+v.Encode(), request.Subdomain), nil)

	if err != nil {
		return ContactResponse{}, err
	}

	req.Header.Add("Authorization", "Bearer "+request.AccessToken)

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return ContactResponse{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 204 {
		var amoErr errors.AmoCRMError
		if err := json.NewDecoder(resp.Body).Decode(&amoErr); err != nil {
			return ContactResponse{}, err
		}
		return ContactResponse{}, fmt.Errorf("error response: %v", amoErr)
	}

	// 204 - контактов нет
	if resp.StatusCode == 204 {
		return ContactResponse{}, fmt.Errorf("empty result")
	}

	var contactResponse ContactResponse
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&contactResponse); err != nil {
		log.Printf("json decode: %v", err)
		return ContactResponse{}, err
	}
	return contactResponse, nil
}
