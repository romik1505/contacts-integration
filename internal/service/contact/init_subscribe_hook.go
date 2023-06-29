package contact

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"week3_docker/internal/client/amo"
	account2 "week3_docker/internal/repository/account"
)

func (cs Service) InitSubscribeHook(ctx context.Context) {
	log.Println("Subscribe hook started")
	defer log.Println("Subscribe hook finished")
	page := 1
	auth := true
	for {
		accounts, err := cs.ar.ListAccounts(ctx, account2.ListAccountFilter{
			Page:          page,
			AmoAuthorized: &auth,
		})
		if errors.Is(err, gorm.ErrRecordNotFound) || len(accounts) == 0 {
			return
		}
		for _, account := range accounts {
			resp, err := cs.amoClient.WebHookContactsSubscribe(ctx, amo.AccountRequest{
				Subdomain:   account.Subdomain,
				AccessToken: account.AccessToken,
			}, account.ID)
			if err != nil {
				log.Printf("Init Subscribe account=%d error: %v\n", account.ID, err)
				continue
			}
			log.Printf("Init Subscribe: account=%d subscribed=%s\n", account.ID, resp.Destination)
		}
		page++
	}
}
