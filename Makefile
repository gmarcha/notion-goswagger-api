DOCKER-COMPOSE := docker compose

.PHONY: all
all: build up

.PHONY: build
build:
	$(DOCKER-COMPOSE) build

.PHONY: up
up:
	$(DOCKER-COMPOSE) up -d --remove-orphans

.PHONY: down
down:
	$(DOCKER-COMPOSE) down

.PHONY: clean
clean:
	$(DOCKER-COMPOSE) down --volumes

.PHONY: gen
gen:
	go generate ./internal/v1/goswagger/restapi
