run:
	go run ./cmd/main.go

generate:
	buf generate

db\:create:
	goose -dir migrations create $(NAME)

local:
	ENVIRONMENT=local go run ./cmd/main.go

#https://www.amocrm.ru/oauth?client_id=0c599009-2fc4-4e6e-b927-c06cca2ab63b&state=1&mode=popup