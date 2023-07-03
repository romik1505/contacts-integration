FROM golang:alpine as builder
LABEL maintainer="amoCRM"
RUN apk update && apk add --no-cache git
WORKDIR /app/
COPY . .
COPY .env .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/main ./cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/worker ./cmd/worker/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/cli ./cmd/cli/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/bin ./bin
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/.env .
COPY --from=builder /app/scripts ./scripts
EXPOSE 8080

CMD ["./scripts/starter.sh", "./bin/cli", "run", "./bin/main"]
