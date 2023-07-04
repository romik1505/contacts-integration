package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"week3_docker/internal/worker"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)

	workerPool := worker.InitializeWorkerPool()
	workerPool.StartWorkers(context.Background())
	<-done

	fmt.Println("worker finished")
}
