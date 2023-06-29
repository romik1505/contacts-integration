package amo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"week3_docker/internal/errors"
)

const (
	authURLMask = "https://%s/oauth2/access_token"
)

// Параметры запроса POST /oauth2/access_token
type AuthRequest struct {
	ClientID     string `json:"client_id,omitempty"`     // ID Интеграции
	ClientSecret string `json:"client_secret,omitempty"` // Секрет интеграции
	GrantType    string `json:"grant_type,omitempty"`    // refresh_token или authorization_code
	Code         string `json:"code,omitempty"`          // Код авторизации(используется 1 раз grant_type=authorization_code)
	RefreshToken string `json:"refresh_token,omitempty"` // Токен обновления (используется для grant_type=refresh_token)
	RedirectURI  string `json:"redirect_uri,omitempty"`  // URI указанный в настройках интеграции
}

// Ответ от AmoCRM /oauth2/access_token
type AuthTokenPair struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func (a Client) AccessToken(ctx context.Context, subdomain string, request AuthRequest) (AuthTokenPair, error) {
	obj, err := json.Marshal(request)
	if err != nil {
		return AuthTokenPair{}, err
	}

	amoURL := fmt.Sprintf(authURLMask, subdomain)
	resp, err := a.httpClient.Post(amoURL, "application/json", bytes.NewBuffer(obj))
	if err != nil {
		return AuthTokenPair{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 204 {
		var amoErr errors.AmoCRMError
		if err = json.NewDecoder(resp.Body).Decode(&amoErr); err != nil {
			return AuthTokenPair{}, err
		}
		return AuthTokenPair{}, fmt.Errorf("integration error: %v", amoErr)
	}

	info := AuthTokenPair{}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return AuthTokenPair{}, err
	}

	return info, nil
}
