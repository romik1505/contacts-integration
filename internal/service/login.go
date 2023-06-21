package service

import (
	"context"
	"fmt"
	"time"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/config"
	"week3_docker/internal/model"
	"week3_docker/internal/store"
)

func (cs ContactService) Login(ctx context.Context, account *model.Account) error {
	request := model.AuthRequest{
		ClientID:     account.IntegrationID.String,
		ClientSecret: config.Config.APISecretKey,
		GrantType:    amo.GrandTypeAccess,
		Code:         account.AuthCode.String,
		RedirectURI:  config.Config.RedirectURI,
	}

	resp, err := cs.amoClient.AccessToken(ctx, *account, request)
	if err != nil {
		return fmt.Errorf("error access token: %v", err)
	}

	account.AccessToken = store.NewNullString(resp.AccessToken)
	account.RefreshToken = store.NewNullString(resp.RefreshToken)
	account.Expires = store.NewNullInt64(time.Now().Unix() + resp.ExpiresIn)

	err = cs.ar.UpdateAccount(ctx, account)
	if err != nil {
		return err
	}

	return nil
}
