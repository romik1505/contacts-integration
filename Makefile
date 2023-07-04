CURRENT_DIR = $(shell pwd)

# generate proto-> grpc, http server
generate:
	buf generate

# local run
run:
	go mod tidy
	google-chrome http://localhost:8080/swagger/index.html & \
	ENVIRONMENT=local ./scripts/starter.sh ./bin/cli run ./bin/main

tests:
	ENVIRONMENT=test ENV_FILE=$(CURRENT_DIR)/.env go test -v ./...

# checkports for docker
ports:
	./scripts/checkport.sh 8080 3307 9090 11300

wire:
	wire ./internal/server
	wire ./internal/worker

build:
	go build -o ./bin/main ./cmd/server/main.go
	go build -o ./bin/cli ./cmd/cli/main.go
	go build -o ./bin/worker ./cmd/worker/main.go

docker:
	./scripts/checkport.sh 8080 3307 9090 11300
	docker-compose build --no-cache
	google-chrome http://localhost:8080/swagger/index.html & \
	docker compose up

lint:
	golangci-lint run ./... --timeout 60s

mocks:
	mockgen -source=./pkg/api/contact_service/contact_service_grpc.pb.go -destination=./pkg/mocks/grpc/contact_service/mock_service.go
	mockgen -source=./internal/repository/account/repository.go -destination=./pkg/mocks/repository/account/mock_account.go
	mockgen -source=./internal/repository/contact/repository.go -destination=./pkg/mocks/repository/contact/mock_contact.go
	mockgen -source=./internal/repository/integration/repository.go -destination=./pkg/mocks/repository/integration/mock_integration.go
	mockgen -source=./internal/service/contact/contact_service.go -destination=./pkg/mocks/service/contact/mock_contact.go
	mockgen -source=./internal/client/amo/amo.go -destination=./pkg/mocks/client/amo/mock_amo.go
	mockgen -source=./internal/client/unisender/unisender.go -destination=./pkg/mocks/client/unisender/mock_unisender.go