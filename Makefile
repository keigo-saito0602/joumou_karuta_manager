# =======================
# Golang CLI Application
# =======================

BINARY_NAME=joumou_karuta_manager
MIGRATION_PATH=assets/migrations
MIGRATE=migrate -source file://$(MIGRATION_PATH) -database "mysql://user:password@tcp(localhost:3306)/joumou_karuta_manager?multiStatements=true"
APP_PORT=8080

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

# ================
# App / CLI
# ================

.PHONY: run
run: ## Run the CLI app directly
	go run main.go

## 再構築
.PHONY: fast-run
fast-run: docker-down docker-volume-clean docker-rebuild docker-up migrate-up

## DBをリセットして再起動
.PHONY: reset
reset: docker-down docker-volume-clean docker-up

.PHONY: serve
serve: ## Run the HTTP server
	go run main.go serve

.PHONY: build
build: ## Build the Go app
	go build -o $(BINARY_NAME) .

.PHONY: clean
clean: ## Clean build artifact
	rm -f $(BINARY_NAME)

# ================
# docker
# ================

.PHONY: docker-up
docker-up:
	docker compose up -d

.PHONY: docker-down
docker-down:
	docker compose down

## Dockerボリュームを削除（データベースの初期化）
.PHONY: docker-volume-clean
docker-volume-clean:
	docker volume rm joumou_karuta_manager_db_data || true

## Dockerイメージを再ビルド
.PHONY: docker-rebuild
docker-rebuild:
	docker compose build --no-cache



# ================
# Migration
# ================

.PHONY: migrate-up
migrate-up: ## Apply all up migrations
	@echo "🚀 Running migration up..."
	$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## Rollback the last migration
	@echo "⏪ Rolling back last migration..."
	$(MIGRATE) down 1

.PHONY: migrate-version
migrate-version: ## Show current migration version
	@$(MIGRATE) version

.PHONY: new-migration
new-migration: ## Create new migration files. Usage: make new-migration NAME=create_users
	@read -p "Enter migration name (snake_case): " NAME; \
	VERSION=$$(date +%Y%m%d%H%M%S); \
	mkdir -p $(MIGRATION_PATH); \
	touch $(MIGRATION_PATH)/$${VERSION}_$${NAME}.up.sql $(MIGRATION_PATH)/$${VERSION}_$${NAME}.down.sql; \
	echo "🆕 Created: $${VERSION}_$${NAME}.up.sql / .down.sql"

# ================
# Swagger / Docs
# ================

.PHONY: swag-init
swag-init: ## Generate Swagger docs
	swag init --parseDependency --parseInternal

.PHONY: swag-open
swag-open: ## Open Swagger UI
	open http://localhost:$(APP_PORT)/swagger/index.html

# ================
# Lint / Test
# ================

.PHONY: lint
lint: ## Run linters
	golangci-lint run

.PHONY: test
test: ## Run unit tests
	go test ./... -v -cover

.PHONY: test-migrate
test-migrate: ## Run migration test (manually write logic if needed)
	go test ./cmd/migrate -v
