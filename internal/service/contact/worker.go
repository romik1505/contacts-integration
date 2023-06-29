package contact

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/beanstalkd/go-beanstalk"
	"log"
	"time"
	"week3_docker/internal/model"
	"week3_docker/internal/queue"
	proto "week3_docker/pkg/api/contact_service"
)

func (w Worker) DoTasks(ctx context.Context) {
	for {
		taskID, body, err := w.Queue.Con.Reserve(5 * time.Second)
		if err != nil {
			if errors.Is(err, beanstalk.ErrTimeout) {
				log.Printf("Worker[%d]: waiting for message\n", w.ID)
				continue
			}
			return
		}
		var task queue.Task
		if err = json.Unmarshal(body, &task); err != nil {
			log.Printf("Worker[%d] unmarshal: %v", w.ID, err)
		}

		log.Printf("Worker[%d]: performing new task %s", w.ID, task.Type)
		start := time.Now()

		switch task.Type {
		case queue.TaskTypePrimarySync:
			data := &proto.ContactSyncRequest{}
			if err := json.Unmarshal(task.Data, data); err != nil {
				log.Printf("Worker[%d]: %v", w.ID, err)
				continue
			}
			if err = w.Service.DoPrimaryContactSync(ctx, data); err != nil {
				log.Printf("Worker[%d]: %v", w.ID, err)
			}
		case queue.TaskTypeAddContacts:
			data := model.ContactActionsTask{}
			if err := json.Unmarshal(task.Data, &data); err != nil {
				log.Printf("Worker[%d]: %v", w.ID, err)
				continue
			}
			data.Type = "add"
			if err := w.Service.DoAddContacts(ctx, data); err != nil {
				log.Printf("Worker[%d]: %v", w.ID, err)
				data.TryNumber++
				if err = w.Queue.PushContactTask(ctx, data, task.Type); err != nil {
					log.Printf("Worker[%d]: DoTask: PushContactTask: %v", w.ID, err)
				}
			}
		case queue.TaskTypeUpdateContacts:
			data := model.ContactActionsTask{}
			if err := json.Unmarshal(task.Data, &data); err != nil {
				log.Printf("Worker[%d]: %v", w.ID, err)
				continue
			}
			data.Type = "update"
			if err := w.Service.DoUpdateContacts(ctx, data); err != nil {
				log.Printf("Worker[%d]: %v", w.ID, err)
				data.TryNumber++
				if err = w.Queue.PushContactTask(ctx, data, task.Type); err != nil {
					log.Printf("Worker[%d]: DoTask: PushContactTask: %v", w.ID, err)
				}
			}
		case queue.TaskTypeDeleteContacts:
			data := model.ContactActionsTask{}
			if err := json.Unmarshal(task.Data, &data); err != nil {
				log.Printf("Worker[%d]: %v", w.ID, err)
				continue
			}
			if err := w.Service.DoDeleteContacts(ctx, data); err != nil {
				log.Printf("Worker[%d]: %v", w.ID, err)
				data.TryNumber++
				if err = w.Queue.PushContactTask(ctx, data, task.Type); err != nil {
					log.Printf("Worker[%d]: DoTask: PushContactTask: %v", w.ID, err)
				}
			}
		}

		log.Printf("Worker[%d]: finished task %s = %v", w.ID, task.Type, time.Now().Sub(start))

		if err := w.Queue.Con.Delete(taskID); err != nil {
			log.Printf("Worker[%d] delete task: %v", w.ID, err)
		}
	}
}
