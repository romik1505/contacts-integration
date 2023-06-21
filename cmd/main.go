package main

import (
	"log"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/repository/account"
	"week3_docker/internal/server"
	"week3_docker/internal/service"
	"week3_docker/internal/store"
	_ "week3_docker/migrations"
)

func main() {
	st, err := store.NewStore()
	if err != nil {
		log.Fatalf("db connection failed %v", err)
	}

	rep := account.NewAccountRepository(st)
	amo := amo.NewAmoClient()
	cs := service.NewContactService(amo, rep)
	go cs.AutoRefreshTokens()

	s := server.NewServer(cs)
	s.Run()
}
