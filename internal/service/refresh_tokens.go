package service

import (
	"context"
	"fmt"
	"log"
	"time"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/config"
	"week3_docker/internal/model"
	"week3_docker/internal/store"
)

func (cs ContactService) AutoRefreshTokens() {
	log.Println("AutoRefreshTokens started")
	defer log.Println("AutoRefreshTokens finished")
	for {
		err := cs.refreshTokens(context.Background())
		if err != nil {
			log.Printf("err refresh tokens: %v\n", err)
		}
		time.Sleep(time.Minute * 20)
	}
}

func (cs ContactService) refreshTokens(ctx context.Context) error {
	accounts, err := cs.ar.ListAccounts(ctx, model.ListAccountFilter{
		Page:        1,
		Limit:       100,
		NeedRefresh: true,
	})
	if err != nil {
		return fmt.Errorf("RefreshTokens: %v", err)
	}
	for _, account := range accounts {
		ctx := context.Background()
		newTokens, err := cs.amoClient.AccessToken(ctx, account, model.AuthRequest{
			ClientID:     account.IntegrationID.String,
			ClientSecret: config.Config.APISecretKey,
			GrantType:    amo.GrantTypeRefresh,
			RefreshToken: account.RefreshToken.String,
			RedirectURI:  config.Config.RedirectURI,
		})
		if err != nil {
			log.Printf("RefreshTokens: error refresh token for account %s: %v", account.ID, err)
			continue
		}
		account.AccessToken = store.NewNullString(newTokens.AccessToken)
		account.RefreshToken = store.NewNullString(newTokens.RefreshToken)
		account.Expires = store.NewNullInt64(time.Now().Unix() + newTokens.ExpiresIn)
		err = cs.ar.UpdateAccount(ctx, &account)
		if err != nil {
			log.Printf("RefreshTokens: error update tokens for account %s: %v", account.ID, err)
		}
	}
	return nil
}
