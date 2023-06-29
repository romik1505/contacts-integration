CURRENT_DIR = $(shell pwd)

run:
	go run ./cmd/server/main.go

generate:
	buf generate

local:
	ENVIRONMENT=local go run ./cmd/server/main.go

tests:
	ENVIRONMENT=local ENV_FILE=$(CURRENT_DIR)/.env go test -v ./...

mocks:
	mockgen -source=./pkg/api/contact_service/contact_service_grpc.pb.go -destination=./pkg/mocks/grpc/contact_service/mock_service.go
	mockgen -source=./internal/repository/account/repository.go -destination=./pkg/mocks/repository/account/mock_account.go
	mockgen -source=./internal/repository/contact/repository.go -destination=./pkg/mocks/repository/contact/mock_contact.go
	mockgen -source=./internal/repository/integration/repository.go -destination=./pkg/mocks/repository/integration/mock_integration.go
	mockgen -source=./internal/service/contact/contact_service.go -destination=./pkg/mocks/service/contact/mock_contact.go
	mockgen -source=./internal/client/amo/amo.go -destination=./pkg/mocks/client/amo/mock_amo.go
	mockgen -source=./internal/client/unisender/unisender.go -destination=./pkg/mocks/client/unisender/mock_unisender.go