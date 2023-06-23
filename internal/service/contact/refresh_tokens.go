package contact

import (
	"context"
	"fmt"
	"log"
	"time"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/config"
	"week3_docker/internal/model"
)

func (cs Service) AutoRefreshTokens() {
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

func (cs Service) refreshTokens(ctx context.Context) error {
	accounts, err := cs.ar.ListAccounts(ctx, model.ListAccountFilter{
		Page:             1,
		Limit:            100,
		NeedRefresh:      true,
		JoinIntegrations: true,
	})

	if err != nil {
		return fmt.Errorf("RefreshTokens: %v", err)
	}
	for _, account := range accounts {
		ctx := context.Background()

		if len(account.Integrations) == 0 {
			log.Printf("account id=%s without integration", account.ID)
			continue
		}

		newTokens, err := cs.amoClient.AccessToken(ctx, account, model.AuthRequest{
			ClientID:     account.Integrations[0].OuterID,
			ClientSecret: config.Config.APISecretKey,
			GrantType:    amo.GrantTypeRefresh,
			RefreshToken: account.RefreshToken,
			RedirectURI:  fmt.Sprintf(redirectUrlMask, config.Config.HostUrl),
		})
		if err != nil {
			log.Printf("RefreshTokens: error refresh token for account %d: %v", account.ID, err)
			continue
		}
		account.AccessToken = newTokens.AccessToken
		account.RefreshToken = newTokens.RefreshToken
		account.Expires = uint64(time.Now().Unix() + newTokens.ExpiresIn)
		err = cs.ar.UpdateAccount(ctx, &account)
		if err != nil {
			log.Printf("RefreshTokens: error update tokens for account %s: %v", account.ID, err)
		}
	}
	return nil
}
