DOCKER_COMPOSE_PATH := ./scripts/docker

# Colors for output
RED := \033[31m
GREEN := \033[32m
YELLOW := \033[33m
BLUE := \033[34m
RESET := \033[0m

export COMPOSE_BAKE=true

.PHONY: help vendor test up up-db up-build down restart logs ps fmt lint

.DEFAULT_GOAL := help

help: ## Show this help message
	@echo "$(BLUE)Available commands:$(RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}'

vendor: ## Run go mod tidy and vendor
	@go mod tidy && go mod vendor

test: ## Run tests
	@echo "$(YELLOW)Running tests...$(RESET)"
	@go test -cover ./...

up: ## Start services with docker compose
	@echo "$(YELLOW)Starting services...$(RESET)"
	@cd $(DOCKER_COMPOSE_PATH) && docker compose --env-file ../../.env up -d
	@echo "$(GREEN)Services started!$(RESET)"

up-db: ## Start database service image only
	@echo "$(YELLOW)Starting database service...$(RESET)"
	@cd $(DOCKER_COMPOSE_PATH) && docker compose --env-file ../../.env up -d mysql migrate
	@echo "$(GREEN)Database service started!$(RESET)"

up-build: ## Start services rebuilding images
	@echo "$(YELLOW)Starting services and rebuilding images...$(RESET)"
	@cd $(DOCKER_COMPOSE_PATH) && docker compose --env-file ../../.env up -d --build
	@echo "$(GREEN)Services started with rebuild!$(RESET)"

down: ## Stop docker compose services
	@echo "$(YELLOW)Stopping services...$(RESET)"
	@cd $(DOCKER_COMPOSE_PATH) && docker compose --env-file ../../.env down
	@echo "$(GREEN)Services stopped!$(RESET)"

restart: ## Restart docker compose services
	@echo "$(YELLOW)Restarting services...$(RESET)"
	@cd $(DOCKER_COMPOSE_PATH) && docker compose --env-file ../../.env restart
	@echo "$(GREEN)Services restarted!$(RESET)"

logs: ## Show docker compose services logs
	@cd $(DOCKER_COMPOSE_PATH) && docker compose --env-file ../../.env logs -f

ps: ## List docker compose services
	@cd $(DOCKER_COMPOSE_PATH) && docker compose --env-file ../../.env ps

fmt: ## Format Go code
	@echo "$(YELLOW)Formatting code...$(RESET)"
	@go fmt ./...
	@echo "$(GREEN)Code formatted!$(RESET)"

lint: ## Run linter
	@echo "$(YELLOW)Running linter...$(RESET)"
	@golangci-lint run
	@echo "$(GREEN)Linting completed!$(RESET)"
