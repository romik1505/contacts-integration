package queue

import (
	"context"
	"encoding/json"
	"log"
	"week3_docker/internal/model"
)

type Task struct {
	Type string          `json:"type,omitempty"`
	Data json.RawMessage `json:"data"`
}

const (
	TaskTypePrimarySync    = "primary_sync"
	TaskTypeAddContacts    = "add_contacts"
	TaskTypeUpdateContacts = "update_contacts"
	TaskTypeDeleteContacts = "delete_contacts"
)

func (q Queue) PushContactTask(ctx context.Context, task model.ContactActionsTask, taskType string) error {
	err := task.Validate()
	if err != nil {
		log.Printf("PushContactTask: %v", err)
		return nil
	}

	jsonData, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = q.PushTask(ctx, Task{
		Type: taskType,
		Data: jsonData,
	})

	if err != nil {
		return err
	}
	return nil
}
