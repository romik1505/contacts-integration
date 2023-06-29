package contact

import (
	"context"
	"fmt"
	"time"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/config"
	"week3_docker/internal/model"
)

const (
	redirectUrlMask = "%s/api/oauth/sign_in"
)

func (cs Service) Login(ctx context.Context, account *model.Account) error {
	if len(account.Integrations) < 1 {
		return fmt.Errorf("ContactService Login: account id=%d without integration", account.ID)
	}
	request := amo.AuthRequest{
		ClientID:     account.Integrations[0].OuterID,
		ClientSecret: config.Config.APISecretKey,
		GrantType:    amo.GrandTypeAccess,
		Code:         account.AuthCode,
		RedirectURI:  fmt.Sprintf(redirectUrlMask, config.Config.HostUrl),
	}

	resp, err := cs.amoClient.AccessToken(ctx, account.Subdomain, request)
	if err != nil {
		return fmt.Errorf("error access token: %v", err)
	}

	account.AccessToken = resp.AccessToken
	account.RefreshToken = resp.RefreshToken
	account.Expires = uint64(time.Now().Unix() + resp.ExpiresIn)

	amoAcc, err := cs.amoClient.Account(ctx, amo.AccountRequest{
		Subdomain:   account.Subdomain,
		AccessToken: account.AccessToken,
	})
	if err != nil {
		return err
	}

	account.ID = amoAcc.ID
	err = cs.ar.CreateAccount(ctx, account)
	if err != nil {
		return err
	}

	return nil
}
