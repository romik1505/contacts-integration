package amo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"week3_docker/internal/errors"
	"week3_docker/internal/model"
)

const (
	authURLMask = "https://%s/oauth2/access_token"
)

func (a Client) AccessToken(ctx context.Context, account model.Account, request model.AuthRequest) (model.AuthTokenPair, error) {
	obj, err := json.Marshal(request)
	if err != nil {
		return model.AuthTokenPair{}, err
	}

	amoURL := fmt.Sprintf(authURLMask, account.Subdomain)
	resp, err := a.httpClient.Post(amoURL, "application/json", bytes.NewBuffer(obj))
	if err != nil {
		return model.AuthTokenPair{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 204 {
		var amoErr errors.AmoCRMError
		if err = json.NewDecoder(resp.Body).Decode(&amoErr); err != nil {
			return model.AuthTokenPair{}, err
		}
		return model.AuthTokenPair{}, fmt.Errorf("integration error: %v", amoErr)
	}

	info := model.AuthTokenPair{}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return model.AuthTokenPair{}, err
	}

	return info, nil
}
