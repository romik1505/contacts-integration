package queue

import (
	"context"
	"encoding/json"
	"github.com/beanstalkd/go-beanstalk"
	"log"
	"time"
	"week3_docker/internal/config"
)

type Queue struct {
	Con *beanstalk.Conn
}

func NewQueue() *Queue {
	cfg := config.Config.BeanstalkdConfig
	con, err := beanstalk.Dial("tcp", cfg.ConnectionString())
	if err != nil {
		log.Fatalf("NewQueue: %v", err)
	}

	return &Queue{
		Con: con,
	}
}

func (q Queue) PushTask(ctx context.Context, t Task) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	_, err = q.Con.Put(data, 1, 0, 120*time.Second)
	if err != nil {
		return err
	}
	return nil
}
