package contact

import (
	"context"
	"log"
	"week3_docker/internal/client/amo"
	account2 "week3_docker/internal/repository/account"
	"week3_docker/internal/store"
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
		if err != nil || len(accounts) == 0 {
			return
		}
		for _, account := range accounts {
			resp, err := cs.amoClient.WebHookContactsSubscribe(ctx, amo.AccountRequest{
				Subdomain:   account.Subdomain,
				AccessToken: account.AccessToken.String,
			}, account.ID)

			if err != nil {
				log.Printf("Init Subscribe account=%d error: %v\n", account.ID, err)

				amoErr, ok := err.(*amo.WebhookSubscribeError)
				if ok && amoErr.Status == 401 {
					account.AccessToken = store.NewNullString("")
					account.RefreshToken = store.NewNullString("")
					account.Expires = store.NewNullInt64(0)
					if err := cs.ar.UpdateAccount(ctx, &account); err != nil {
						log.Printf("InitSubscribeHook update: %v", err)
					}
				}

				continue
			}
			log.Printf("Init Subscribe: account=%d subscribed=%s\n", account.ID, resp.Destination)
		}
		page++
	}
}
