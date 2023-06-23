package main

import (
	"log"
	"week3_docker/internal/client/amo"
	"week3_docker/internal/client/unisender"
	"week3_docker/internal/repository/account"
	contact_repository "week3_docker/internal/repository/contact"
	"week3_docker/internal/repository/integration"
	"week3_docker/internal/server"
	"week3_docker/internal/service/contact"
	"week3_docker/internal/store"
)

func main() {
	st, err := store.NewStore()
	if err != nil {
		log.Fatalf("db connection failed %v", err)
	}

	ar := account.NewAccountRepository(st)
	cr := contact_repository.NewRepository(st)
	ir := integration.NewRepository(st)
	amo := amo.NewAmoClient()
	uni := unisender.NewClient()
	cs := contact.NewService(amo, uni, ar, cr, ir)
	go cs.AutoRefreshTokens()

	s := server.NewServer(cs)
	s.Run()
}
