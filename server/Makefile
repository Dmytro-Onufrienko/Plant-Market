MIGRATE    = ~/go/bin/migrate
AIR        = ~/go/bin/air
MIGRATIONS = internal/db/migrations
DB_URL    ?= $(shell grep DATABASE_URL .env | cut -d '=' -f2-)

.PHONY: dev docker-up docker-down migrate-up migrate-down build test

dev:
	$(AIR) -c .air.toml

docker-up:
	docker compose -f docker/docker-compose.yml up -d

docker-down:
	docker compose -f docker/docker-compose.yml down

migrate-up:
	$(MIGRATE) -path $(MIGRATIONS) -database "$(DB_URL)" up

migrate-down:
	$(MIGRATE) -path $(MIGRATIONS) -database "$(DB_URL)" down 1

build:
	go build -o bin/server ./cmd/server

test:
	go test ./...
