.phony: build

# Alias for docker-compose
COMPOSE := docker-compose

up: ## Boot all containers
	$(COMPOSE) up -d

build:
	$(COMPOSE) build --no-cache

down:
	$(COMPOSE) kill

build-one:
	$(COMPOSE) build --no-cache $(tag)
