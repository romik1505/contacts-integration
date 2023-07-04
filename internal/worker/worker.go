package worker

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/beanstalkd/go-beanstalk"
	"log"
	"time"
	"week3_docker/internal/model"
	"week3_docker/internal/queue"
	"week3_docker/internal/service/contact"
	proto "week3_docker/pkg/api/contact_service"
)

type WorkerPool struct {
	Queue   *queue.Queue
	Service contact.IService
	Num     int
}

func NewWorkerPool(q *queue.Queue, s *contact.Service) WorkerPool {
	w := WorkerPool{
		Queue:   q,
		Service: s,
		Num:     2,
	}

	return w
}

func (w WorkerPool) DoTasks(ctx context.Context, id int) {
	for {
		taskID, body, err := w.Queue.Con.Reserve(2 * time.Second)
		if err != nil {
			if errors.Is(err, beanstalk.ErrTimeout) {
				log.Printf("Worker[%d]: waiting for message\n", id)
				continue
			}
			return
		}
		var task queue.Task
		if err = json.Unmarshal(body, &task); err != nil {
			log.Printf("Worker[%d] unmarshal: %v", id, err)
		}

		log.Printf("Worker[%d]: performing new task %s", id, task.Type)
		start := time.Now()

		switch task.Type {
		case queue.TaskTypePrimarySync:
			data := &proto.ContactSyncRequest{}
			if err := json.Unmarshal(task.Data, data); err != nil {
				log.Printf("Worker[%d]: %v", id, err)
				continue
			}
			if err = w.Service.DoPrimaryContactSync(ctx, data); err != nil {
				log.Printf("Worker[%d]: %v", id, err)
			}
		case queue.TaskTypeAddContacts:
			data := model.ContactActionsTask{}
			if err := json.Unmarshal(task.Data, &data); err != nil {
				log.Printf("Worker[%d]: %v", id, err)
				continue
			}
			data.Type = "add"
			if err := w.Service.DoAddContacts(ctx, data); err != nil {
				log.Printf("Worker[%d]: %v", id, err)
				data.TryNumber++
				if err = w.Queue.PushContactTask(ctx, data, task.Type); err != nil {
					log.Printf("Worker[%d]: DoTask: PushContactTask: %v", id, err)
				}
			}
		case queue.TaskTypeUpdateContacts:
			data := model.ContactActionsTask{}
			if err := json.Unmarshal(task.Data, &data); err != nil {
				log.Printf("Worker[%d]: %v", id, err)
				continue
			}
			data.Type = "update"
			if err := w.Service.DoUpdateContacts(ctx, data); err != nil {
				log.Printf("Worker[%d]: %v", id, err)
				data.TryNumber++
				if err = w.Queue.PushContactTask(ctx, data, task.Type); err != nil {
					log.Printf("Worker[%d]: DoTask: PushContactTask: %v", id, err)
				}
			}
		case queue.TaskTypeDeleteContacts:
			data := model.ContactActionsTask{}
			if err := json.Unmarshal(task.Data, &data); err != nil {
				log.Printf("Worker[%d]: %v", id, err)
				continue
			}
			if err := w.Service.DoDeleteContacts(ctx, data); err != nil {
				log.Printf("Worker[%d]: %v", id, err)
				data.TryNumber++
				if err = w.Queue.PushContactTask(ctx, data, task.Type); err != nil {
					log.Printf("Worker[%d]: DoTask: PushContactTask: %v", id, err)
				}
			}
		}

		log.Printf("Worker[%d]: finished task %s = %v", id, task.Type, time.Since(start))

		if err := w.Queue.Con.Delete(taskID); err != nil {
			log.Printf("Worker[%d] delete task: %v", id, err)
		}
	}
}

func (w WorkerPool) StartWorkers(ctx context.Context) {
	for i := 0; i < w.Num; i++ {
		go w.DoTasks(ctx, i+1)
	}
}
