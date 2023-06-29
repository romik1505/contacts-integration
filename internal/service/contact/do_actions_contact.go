package contact

import (
	"context"
	"log"
	"week3_docker/internal/client/unisender"
	"week3_docker/internal/mapper"
	"week3_docker/internal/model"
	contact2 "week3_docker/internal/repository/contact"
)

func (s Service) DoAddContacts(ctx context.Context, task model.ContactActionsTask) error {
	if task.TryNumber > 2 {
		log.Printf("DoAddContacts: task type=%s, tries %d", task.Type, task.TryNumber)
		return nil
	}

	_, err := s.cr.InsertContacts(ctx, task.Contacts)
	if err != nil {
		return err
	}
	req, err := mapper.ConvertContactsToUnisenderRequestParams(
		task.Contacts,
		task.UnisenderKey,
		task.UnisenderListID,
	)
	var res unisender.ImportContactsResponse
	if err == nil {
		res, err = s.uniClient.ImportContacts(ctx, req)
		if err != nil {
			return err
		}
	}

	mp := unisender.LogsToMap(res.Result)

	for i, contact := range task.Contacts {
		reason, exist := mp[i]
		if !exist { // success import
			contact.Sync = true
		} else { // addition failed
			contact.ReasonOutSync = reason
		}
		if err = s.cr.UpdateContact(ctx, &contact); err != nil {
			log.Printf("DoAddContacts: update by logs: %v", err)
		}
	}

	return nil
}

func (s Service) DoUpdateContacts(ctx context.Context, task model.ContactActionsTask) error {
	if task.TryNumber > 2 {
		log.Printf("DoUpdateContacts: task type=%s, tries %d", task.Type, task.TryNumber)
		return nil
	}

	if err := s.DoDeleteContacts(ctx, task); err != nil {
		return err
	}

	if err := s.DoAddContacts(ctx, task); err != nil {
		return err
	}

	return nil
}

func (s Service) DoDeleteContacts(ctx context.Context, task model.ContactActionsTask) error {
	if task.TryNumber > 2 {
		log.Printf("DoDeleteContacts: task type=%s, tries %d", task.Type, task.TryNumber)
		return nil
	}

	err := s.cr.UpdateContactsByAmoIDs(ctx, task.IDs, &model.Contact{
		Type: "delete",
		Sync: false,
	})
	if err != nil {
		return err
	}

	f := false
	contacts, err := s.cr.ListContacts(ctx, contact2.ListContactsFilter{
		AccountID: task.AccountID,
		Page:      1,
		Type:      "delete",
		AmoIDs:    task.IDs,
		Sync:      &f,
	})
	if err != nil {
		return err
	}

	var res unisender.ImportContactsResponse
	req, err := mapper.ConvertContactsToUnisenderRequestParams(contacts, task.UnisenderKey, task.UnisenderListID)
	if err == nil {
		res, err = s.uniClient.ImportContacts(ctx, req)
		if err != nil {
			return err
		}
	}

	mp := unisender.LogsToMap(res.Result)

	for i, contact := range contacts {
		reason, exist := mp[i]
		if !exist || contact.ReasonOutSync != "" { // success delete
			// hard delete
			if err = s.cr.DeleteContact(ctx, &contact); err != nil {
				log.Printf("DoDeleteContacts: delete by logs: %v", err)
			}
		} else {
			contact.Sync = false
			contact.ReasonOutSync = reason
			if err := s.cr.UpdateContact(ctx, &contact); err != nil {
				log.Printf("DoDeleteContacts: reason: %v", err)
			}
		}
	}

	return nil
}
