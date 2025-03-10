export $(shell cat .env | xargs)

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

run:
	go run cmd/main.go

build:
	go build -o monitoring-app cmd/main.go

test:
	go test -race -count 100 ./...

migrate-up:
	goose -dir sql/migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir sql/migrations postgres "$(DB_URL)" down
