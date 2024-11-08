GOOSE_DRIVER ?= postgres
GOOSE_DBSTRING ?= "postgresql://postgres:Postgres123@45.149.207.141:55432/gopress"
GOOSE_MIGRATIONS_PATH ?= "./migrations"

migrate-up:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_PATH) up
migrate-status:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_PATH) status
migrate-down:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_PATH) down
migrate-reset:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_PATH) reset
run:
	go run cmd/server/main.go
dev:
	nodemon
	
sqlgen:
	sqlc generate

.PHONY: migrate-up migrate-status migrate-down migrate-reset run dev sqlgen