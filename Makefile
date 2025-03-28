include .local_env

.PHONY: install-migrate
install-migrate:
	@chmod +x ./scripts/install_migrate.sh && ./scripts/install_migrate.sh

.PHONY: setup-dev
setup-dev: install-migrate
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
	go install github.com/golang/mock/mockgen@v1.5.0

.PHONY: mock-gen
mock-gen:
	go generate ./...

.PHONY: lint
lint:
	@ golangci-lint run ./...

.PHONY: format
format:
	go fmt ./...

.PHONY: unit-test
unit-test:
	go test ./...

set_dot_env:
	@set -o allexport; source .env; set +o allexport;

.PHONY: create_migration
create-migration:
	migrate create -ext=sql -dir=migrations/postgres -seq init

.PHONY: migrate-up
migrate-up:
	migrate -path=migrations/postgres -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path=migrations/postgres -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down
