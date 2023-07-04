package contact

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/client/unisender"
	"week3_docker/internal/mapper"
	"week3_docker/internal/model"
	"week3_docker/internal/queue"
	"week3_docker/internal/repository/contact"
	"week3_docker/internal/schemas"
	proto "week3_docker/pkg/api/contact_service"
)

func (s Service) PrimaryContactsSync(ctx context.Context, req schemas.ContactSyncRequest) error {
	fmt.Printf("id=%d, unisender_key=%s\n", req.AccountID, req.UnisenderKey)
	data, err := json.Marshal(&proto.ContactSyncRequest{
		AccountId:    req.AccountID,
		UnisenderKey: req.UnisenderKey,
	})
	if err != nil {
		return err
	}
	task := queue.Task{
		Type: queue.TaskTypePrimarySync,
		Data: data,
	}
	if err = s.queue.PushTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (s Service) DoPrimaryContactSync(ctx context.Context, req *proto.ContactSyncRequest) error {
	account, err := s.ar.GetAccount(ctx, req.GetAccountId())
	if err != nil {
		return err
	}

	resp, err := s.amoClient.WebHookContactsSubscribe(ctx, amo.AccountRequest{
		Subdomain:   account.Subdomain,
		AccessToken: account.AccessToken.String,
	}, account.ID)
	if err != nil {
		return err
	}
	log.Printf("webhook created %v\n", resp)

	listsResp, err := s.uniClient.GetLists(ctx, req.GetUnisenderKey())
	if err != nil {
		if err.Error() == "invalid_api_key" {
			log.Printf("invalid_api_key for account %d", req.GetAccountId())
			err := s.ar.UpdateAccount(ctx, &model.Account{ID: req.GetAccountId(), UnisenderKey: ""})
			if err != nil {
				return err
			}
		}
		return err
	}

	var unisenderListID uint64
	for _, v := range listsResp.Result {
		if v.Title == "amoCRM" {
			unisenderListID = v.ID
			break
		}
	}
	if unisenderListID == 0 {
		unisenderListID, err = s.uniClient.CreateList(ctx, req.GetUnisenderKey(), "amoCRM")
		if err != nil {
			return err
		}
	}

	account.UnisenderKey = req.GetUnisenderKey()
	account.UnisenderListID = unisenderListID
	err = s.ar.UpdateAccount(ctx, &account)
	if err != nil {
		return err
	}

	params := amo.ContactQueryParams{
		Page:  1,
		Limit: 2,
	}

	for {
		contacts, err := s.amoClient.ListContacts(context.Background(), amo.AccountRequest{
			Subdomain:   account.Subdomain,
			AccessToken: account.AccessToken.String,
		}, params)
		if err != nil {
			break
		}

		cs := mapper.ConvertAmoContactsToModel(contacts.Embedded.Contacts, model.ContactTypePrimarySync)
		if len(cs) == 0 {
			params.Page++
			continue
		}
		_, err = s.cr.InsertContacts(context.Background(), cs)
		if err != nil {
			log.Printf("PrimaryContactSync: insert error %v", err)
		}
		params.Page++
	}

	if account.UnisenderListID == 0 {
		listID, err := s.uniClient.CreateList(ctx, account.UnisenderKey, "amoCRM")
		if err != nil {
			return err
		}
		account.UnisenderListID = listID
		if err := s.ar.UpdateAccount(ctx, &account); err != nil {
			return err
		}
	}

	page := 1

	for {
		b := false
		contacts, err := s.cr.ListContacts(ctx, contact.ListContactsFilter{
			AccountID: account.ID,
			Page:      page,
			Limit:     100,
			Type:      model.ContactTypePrimarySync,
			Sync:      &b,
		})
		if err != nil || len(contacts) == 0 {
			break
		}

		var resp unisender.ImportContactsResponse
		req, err := mapper.ConvertContactsToUnisenderRequestParams(contacts, account.UnisenderKey, account.UnisenderListID)
		if err == nil {
			resp, err = s.uniClient.ImportContacts(ctx, req)
			if err != nil {
				return err
			}
		}

		mp := unisender.LogsToMap(resp.Result)

		for i, cont := range contacts {
			reason, exist := mp[i]
			if !exist { // success import
				cont.Sync = true
			} else {
				cont.Sync = false
				cont.ReasonOutSync = reason
			}
			if err := s.cr.UpdateContact(ctx, &cont); err != nil {
				log.Println(err.Error())
			}
		}

		page++
	}
	return nil
}
