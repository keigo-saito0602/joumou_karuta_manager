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

# =======================
# Golang CLI Application
# =======================

BINARY_NAME=joumou_karuta_manager
MIGRATION_PATH=assets/migrations
APP_PORT=8080

# Makefile 例
MIGRATE=docker run --rm \
  --network joumou_karuta_manager_default \
  -v "$(shell pwd)/assets/migrations:/migrations" \
  migrate/migrate \
  -source=file:///migrations \
  -database "$(DATABASE_URL)"

# =======================
# Utility
# =======================

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

.PHONY: check-env
check-env: ## Check required env vars are set
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
init-env: ## Create .env from example if not exists
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
run: ## Run the CLI app directly
	go run main.go

.PHONY: serve
serve: ## Run the HTTP server
	go run main.go serve

.PHONY: build
build: ## Build the Go app
	go build -o $(BINARY_NAME) .

.PHONY: clean
clean: ## Clean build artifact
	rm -f $(BINARY_NAME)

# =======================
# Docker
# =======================

.PHONY: docker-up
docker-up: ## Start docker containers
	docker compose up -d

.PHONY: docker-down
docker-down: ## Stop docker containers
	docker compose down

.PHONY: docker-volume-clean
docker-volume-clean: ## Remove DB volume
	docker volume rm joumou_karuta_manager_db_data || true

.PHONY: docker-rebuild
docker-rebuild: ## Rebuild docker containers
	docker compose build --no-cache

.PHONY: fast-run
fast-run: docker-down docker-volume-clean docker-rebuild docker-up migrate-up ## Rebuild all and migrate

.PHONY: reset
reset: docker-down docker-volume-clean docker-up ## Reset DB and restart

# =======================
# Migration
# =======================

.PHONY: migrate-up
migrate-up: check-env ## Apply all up migrations
	@echo "🚀 Running migration up..."
	$(MIGRATE) up

.PHONY: migrate-down
migrate-down: check-env ## Rollback the last migration
	@echo "⏪ Rolling back last migration..."
	$(MIGRATE) down 1

.PHONY: migrate-version
migrate-version: check-env ## Show current migration version
	@$(MIGRATE) version

.PHONY: new-migration
new-migration: ## Create new migration files. Usage: make new-migration NAME=create_users
	@read -p "Enter migration name (snake_case): " NAME; \
	VERSION=$$(date +%Y%m%d%H%M%S); \
	mkdir -p $(MIGRATION_PATH); \
	touch $(MIGRATION_PATH)/$${VERSION}_$${NAME}.up.sql $(MIGRATION_PATH)/$${VERSION}_$${NAME}.down.sql; \
	echo "🆕 Created: $${VERSION}_$${NAME}.up.sql / .down.sql"

# =======================
# Swagger
# =======================

.PHONY: swag-init
swag-init: ## Generate Swagger docs
	swag init --parseDependency --parseInternal

.PHONY: swag-open
swag-open: ## Open Swagger UI
	open http://localhost:$(APP_PORT)/swagger/index.html

# =======================
# Lint / Test
# =======================

.PHONY: lint
lint: ## Run linters
	golangci-lint run

.PHONY: test
test: ## Run unit tests
	go test ./... -v -cover

.PHONY: test-migrate
test-migrate: ## Run migration tests
	go test ./cmd/migrate -v
