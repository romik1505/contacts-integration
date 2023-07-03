package contact

import (
	"context"
	"fmt"
	"week3_docker/internal/mapper"
	"week3_docker/internal/model"
	"week3_docker/internal/queue"
	"week3_docker/internal/schemas"
)

func (s Service) ContactActionsHook(ctx context.Context, req schemas.ContactActionsHookRequest) error {
	account, err := s.ar.GetAccount(ctx, req.ID)
	if err != nil {
		return err
	}
	if account.UnisenderKey == "" || account.UnisenderListID == 0 {
		return fmt.Errorf("require primary sync again")
	}

	addContacts, addIDs := mapper.ConvertAmoContactsWithIDs(req.Contacts.Add, req.ID, model.ContactTypeAdd)
	updateContacts, updateIDs := mapper.ConvertAmoContactsWithIDs(req.Contacts.Update, req.ID, model.ContactTypeUpdate)
	deleteIDs := mapper.AmoContactsIDs(req.Contacts.Delete)

	taskData := model.ContactActionsTask{
		AccountID:       account.ID,
		UnisenderKey:    account.UnisenderKey,
		UnisenderListID: account.UnisenderListID,
	}

	taskData.Contacts = addContacts
	taskData.IDs = addIDs
	taskData.Type = model.ContactTypeAdd
	err = s.queue.PushContactTask(ctx, taskData, queue.TaskTypeAddContacts)
	if err != nil {
		return err
	}

	taskData.Contacts = updateContacts
	taskData.IDs = updateIDs
	taskData.Type = model.ContactTypeUpdate
	err = s.queue.PushContactTask(ctx, taskData, queue.TaskTypeUpdateContacts)
	if err != nil {
		return err
	}

	taskData.Contacts = nil
	taskData.IDs = deleteIDs
	taskData.Type = model.ContactTypeDelete
	err = s.queue.PushContactTask(ctx, taskData, queue.TaskTypeDeleteContacts)
	if err != nil {
		return err
	}

	return nil
}
