ifneq (,$(wildcard ./.env))
    include .env
    export
endif

POSTGRES_DSN = "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}"

PROJECT_DIR = $(CURDIR)
PROJECT_BIN = ${PROJECT_DIR}/bin

# Deps versions
WIRE_VERSION = "0.6.0"
GOLANGCI_VERSION = "2.4.0"
GOOSE_VERSION = "3.24.3"

help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
.PHONY: help

deps-bin:  ## Install development dependencies and tools
	echo "Install dependencies"
	@GOBIN=${PROJECT_BIN} go install github.com/google/wire/cmd/wire@v${WIRE_VERSION}
	@GOBIN=${PROJECT_BIN} go install github.com/pressly/goose/v$(shell echo $(GOOSE_VERSION) | cut -d. -f1)/cmd/goose@v${GOOSE_VERSION}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${PROJECT_BIN} v${GOLANGCI_VERSION}
.PHONY: deps-bin

lint: ## Run linter
	@${PROJECT_BIN}/golangci-lint run --config=./.golangci.yml
.PHONY: lint

dotenv: ## Generate .env file from .env.example
	@cp .env.example .env
.PHONY: dotenv

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

up: ## Up service and all necessary infrastructure
	@docker-compose up -d --build
	@ docker-compose logs -f
.PHONY: up

down: ## Down service and all necessary infrastructure
	@docker-compose down
.PHONY: down

rm: ## Remove service and all necessary infrastructure
	@docker-compose down -v
.PHONY: infra-rm

run: ## Build and run application (go run) (infrastructure should exits)
	@go run ./cmd/main.go
.PHONY: run