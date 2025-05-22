# =======================
# Load .env file if exists
# =======================
ifneq (,$(wildcard .env))
	include .env
	export
endif

# =======================
# Environment Variables
# =======================

DB_HOST ?= db
DB_PORT ?= 3306
DB_USER ?= user
DB_PASSWORD ?= password
DB_NAME ?= joumou_karuta_manager

PROJECT_NAME ?= joumou_karuta_manager
DOCKER_NETWORK_NAME ?= joumou_karuta_manager_default

# =======================
# Golang CLI Application
# =======================

BINARY_NAME ?= $(PROJECT_NAME)
MIGRATION_PATH ?= assets/migrations
APP_PORT ?= 8080

# Makefile 例
MIGRATE=docker run --rm \
  --network $(DOCKER_NETWORK_NAME) \
  -v "$(shell pwd)/$(MIGRATION_PATH):/migrations" \
  migrate/migrate \
  -source=file:///migrations \
  -database "$(DATABASE_URL)"

# =======================
# Utility
# =======================

.PHONY: help
help: ## [make help] Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

.PHONY: check-env
check-env: ## [make check-env] Check required env vars are set
	@missing_vars=0; \
	for var in DB_HOST DB_PORT DB_USER DB_PASSWORD DB_NAME; do \
		if [ -z "$${!var}" ]; then \
			echo "❌ Environment variable $$var is not set"; \
			missing_vars=1; \
		fi \
	done; \
	if [ $$missing_vars -eq 1 ]; then \
		echo "💡 You can copy from .env.example: cp .env.example .env"; \
		exit 1; \
	fi

.PHONY: init-env
init-env: ## [make init-env] Create .env from example if not exists
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "✅ .env created from .env.example"; \
	else \
		echo "📝 .env already exists"; \
	fi

# =======================
# Application / CLI
# =======================

.PHONY: run
run: ## [make run] Run the CLI app directly
	go run main.go

.PHONY: serve
serve: ## [make serve] Run the HTTP server
	go run main.go serve

.PHONY: build
build: ## [make build] Build the Go app
	go build -o $(BINARY_NAME) .

.PHONY: clean
clean: ## [make clean] Clean build artifact
	rm -f $(BINARY_NAME)

# =======================
# Docker
# =======================

.PHONY: docker-up
docker-up: ## [make docker-up] Start docker containers
	docker compose up -d

.PHONY: docker-down
docker-down: ## [make docker-down] Stop docker containers
	docker compose down

.PHONY: docker-volume-clean
docker-volume-clean: ## [make docker-volume-clean] Remove DB volume
	docker volume rm $(DB_VOLUME_NAME) || true

.PHONY: docker-rebuild
docker-rebuild: ## [make docker-rebuild] Rebuild docker containers
	docker compose build --no-cache

.PHONY: docker-exec
docker-exec: ## [make docker-exec] Exec docker containers
	docker exec -it $(PROJECT_NAME) sh

.PHONY: fast-run
fast-run: reset swag-init docker-up migrate-up logs ## [make fast-run] Rebuild docker server

.PHONY: launch
launch: swag-init docker-up logs ## [make launch] Start server without rebuilding

.PHONY: reset
reset: docker-down docker-volume-clean docker-rebuild ## [make reset] Reset DB and restart

.PHONY: logs
logs: ## [make logs] Follow app logs
	docker logs -f $(PROJECT_NAME)

# =======================
# Migration
# =======================

.PHONY: migrate-up
migrate-up: check-env ## [make migrate-up] Apply all up migrations
	@echo "🚀 Running migration up..."
	$(MIGRATE) up

.PHONY: migrate-down
migrate-down: check-env ## [make migrate-down] Rollback the last migration
	@echo "⏪ Rolling back last migration..."
	$(MIGRATE) down 1

.PHONY: migrate-version
migrate-version: check-env ## [make migrate-version] Show current migration version
	@$(MIGRATE) version

.PHONY: new-migration
new-migration: ## [make new-migration] Create new migration files. Usage: make new-migration NAME=create_users
	@read -p "Enter migration name (snake_case): " NAME; \
	VERSION=$$(date +%Y%m%d%H%M%S); \
	mkdir -p $(MIGRATION_PATH); \
	touch $(MIGRATION_PATH)/$${VERSION}_$${NAME}.up.sql $(MIGRATION_PATH)/$${VERSION}_$${NAME}.down.sql; \
	echo "🆕 Created: $${VERSION}_$${NAME}.up.sql / .down.sql"

# =======================
# Swagger
# =======================

.PHONY: swag-init
swag-init: ## [make swag-init] Generate Swagger docs
	swag init --parseDependency --parseInternal

.PHONY: swag-open
swag-open: ## [make swag-open] Open Swagger UI
	open http://localhost:$(APP_PORT)/swagger/index.html

# =======================
# Lint / Test
# =======================

.PHONY: lint
lint: ## [make lint] Run linters
	golangci-lint run --concurrency=2

.PHONY: test
test: ## [make test] Run unit tests
	go test ./... -v -cover

.PHONY: test-migrate
test-migrate: ## [make test-migrate] Run migration tests
	go test ./cmd/migrate -v

.PHONY: generate-mock
generate-mock: ## [make generate-mock] 任意の usecase ファイルからモックを生成する 例: make generate-mock USECASE=usecase/user_usecase.go
	@echo "🔧 Generating mock for $(USECASE)..."
	@test -n "$(USECASE)" || (echo "❌ USECASEパラメータが必要です（例: make generate-mock USECASE=usecase/user_usecase.go）" && exit 1)
	@USECASE_FILE=$(USECASE) && \
	BASENAME=$$(basename $$USECASE_FILE .go) && \
	MOCK_PATH=interface/handler/mocks/$${BASENAME}_mock.go && \
	mockgen -source=$$USECASE_FILE -destination=$$MOCK_PATH -package=mocks && \
	echo "✅ Mock generated: $$MOCK_PATH"

.PHONY: cover
cover:
	go test -coverprofile=coverage.out \
		./interface/handler/... \
		./usecase/... \
		./infrastructure/repository/... \
		./validation/... \
		./auth/... \
		./util/...
	go tool cover -html=coverage.out
