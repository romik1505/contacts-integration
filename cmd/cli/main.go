package main

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const workerProcess = "worker"

func RunWorker(context *cli.Context) (*exec.Cmd, error) {
	workerFile := os.Getenv("WORKER_FILE")
	if workerFile == "" {
		workerFile = "./bin/worker"
	}

	cmd := exec.Command(workerFile)

	err := cmd.Start()

	if err != nil {
		fmt.Printf("RunWorker: %v", err)
		return nil, err
	}

	return cmd, nil
}

func RunServer(context *cli.Context) (*exec.Cmd, error) {
	workerFile := os.Getenv("SERVER_FILE")
	if workerFile == "" {
		workerFile = "./bin/main"
	}

	cmd := exec.Command(workerFile)

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	fmt.Println("started server")

	return cmd, nil
}

func Stop(context *cli.Context) error {
	killWorkerProcess()
	return nil
}

func main() {
	app := &cli.App{
		Name:     "worker",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "start worker process",
				Action: func(ctx *cli.Context) error {
					_, err := RunWorker(ctx)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "start only server without worker",
				Action: func(ctx *cli.Context) error {
					cmd, err := RunServer(ctx)
					defer fmt.Println("finished server")

					if err != nil {
						return err
					}
					err = cmd.Wait()
					if err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:    "docker",
				Aliases: []string{"d"},
				Usage:   "run workers and server inside docker",
				Action: func(ctx *cli.Context) error {
					_, err := RunWorker(ctx)
					if err != nil {
						return err
					}
					s, err := RunServer(ctx)
					if err != nil {
						return err
					}
					err = s.Wait()
					if err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:    "stop",
				Usage:   "stop worker process",
				Aliases: []string{"st"},
				Action:  Stop,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error run app: %v\n", err)
		return
	}
}

func findAndKillProcess(path string, info os.FileInfo, err error) error {
	if err != nil {
		// log.Println(err)
		return nil
	}

	if strings.Count(path, "/") == 3 {
		if strings.Contains(path, "/status") {
			pid, err := strconv.Atoi(path[6:strings.LastIndex(path, "/")])
			if err != nil {
				log.Println(err)
				return nil
			}

			f, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println(err)
				return nil
			}

			name := string(f[6:bytes.IndexByte(f, '\n')])

			if name == workerProcess {
				fmt.Printf("PID: %d, Name: %s will be killed.\n", pid, name)

				proc, err := os.FindProcess(pid)
				if err != nil {
					log.Println(err.Error())
				}

				if err = proc.Kill(); err != nil {
					log.Println(err.Error())
				}

				return io.EOF
			}
		}
	}

	return nil
}

func killWorkerProcess() {
	if err := filepath.Walk("/proc", findAndKillProcess); err != nil {
		log.Printf("error walking the path /proc")
	}
}
