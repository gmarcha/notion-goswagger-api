DOCKER-COMPOSE := docker compose
DOCKER-COMPOSE-PATH := docker-compose.yaml

.PHONY: all
all: build up

.PHONY: build
build:
	$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) build

.PHONY: up
up:
	$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) up -d --remove-orphans

.PHONY: down
down:
	$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) down

.PHONY: clean
clean:
	$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) down --volumes

.PHONY: gen
gen:
	go generate ./internal/v1/goswagger/restapi
