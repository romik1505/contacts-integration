package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"week3_docker/internal/worker"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Kill, os.Interrupt)

	workerPool := worker.InitializeWorkerPool()
	workerPool.StartWorkers(context.Background())
	<-done

	fmt.Println("worker finished")
}
