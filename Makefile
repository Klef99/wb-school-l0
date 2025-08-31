ifneq (,$(wildcard ./.env))
    include .env
    export
endif

POSTGRES_DSN = "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}"

PROJECT_DIR = $(CURDIR)
PROJECT_BIN = ${PROJECT_DIR}/bin

# Deps versions
MOCKERY_BRANCH = "v2@v2.53.4"
WIRE_VERSION = "0.6.0"
GOLANGCI_VERSION = "2.4.0"
GOOSE_VERSION = "3.24.3"

help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
.PHONY: help

deps-bin:  ## Install service dependencies
	@GOBIN=${PROJECT_BIN} go install github.com/google/wire/cmd/wire@v${WIRE_VERSION}
	@GOBIN=${PROJECT_BIN} go install github.com/vektra/mockery/${MOCKERY_BRANCH}
	@GOBIN=${PROJECT_BIN} go install github.com/pressly/goose/v$(shell echo $(GOOSE_VERSION) | cut -d. -f1)/cmd/goose@v${GOOSE_VERSION}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${PROJECT_BIN} v${GOLANGCI_VERSION}
.PHONY: deps-bin

lint: ## Run linter
	@${PROJECT_BIN}/golangci-lint run --config=./.golangci.yml
.PHONY: lint

gen-wire: ## Generate wire_gen.go file
	@${PROJECT_BIN}/wire gen github.com/klef99/wb-school-l0/cmd/di
.PHONY: gen-wire

migrations: ## Run migrations utility
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${POSTGRES_DSN} GOOSE_MIGRATION_DIR="./migrations" ${PROJECT_BIN}/goose $(GOOSE_CMD)
.PHONY: migrations

migrations-up:
	@GOOSE_CMD="up" make migrations ## Up migrations to latest version
.PHONY: migrations-up

migrations-down:
	@GOOSE_CMD="down" make migrations ## Down migrations to latest version
.PHONY: migrations-down

infra-up: ## Up infrastructure for service
	@docker-compose up -d --build
	@ docker-compose logs -f
.PHONY: infra-up

infra-down: ## Down infrastructure for service
	@docker-compose down
.PHONY: infra-down

infra-rm: ## Remove infrastructure for service
	@docker-compose down -v
.PHONY: infra-rm

run: gen-wire ## Build and run application (go run)
	@go run ./cmd/main.go
.PHONY: run