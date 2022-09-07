DOCKER-COMPOSE := docker compose
PATH := docker-compose.yaml

.PHONY: all
all: build up

.PHONY: build
build:
	$(DOCKER-COMPOSE) -f $(PATH) build

.PHONY: up
up:
	$(DOCKER-COMPOSE) -f $(PATH) up -d --remove-orphans

.PHONY: down
down:
	$(DOCKER-COMPOSE) -f $(PATH) down

.PHONY: clean
clean:
	$(DOCKER-COMPOSE) -f $(PATH) down --volumes

.PHONY: gen
gen:
	go generate ./internal/v1/goswagger/restapi
