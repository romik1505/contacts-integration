package main

import "week3_docker/internal/server"

func main() {
	s := server.InitializeServer()
	s.Run()
}
