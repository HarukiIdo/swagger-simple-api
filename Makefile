COMMAND:=
DOCKER_COMPOSE_VERSION_CHECKER := $(shell docker compose > /dev/null 2>&1 ; echo $$?)
ifeq ($(DOCKER_COMPOSE_VERSION_CHECKER), 0)
	DOCKER_COMPOSE_IMPL=docker compose
else
	DOCKER_COMPOSE_IMPL=docker-compose
endif

.PHONY: docker-compose/command
docker-compose/command:
	$(DOCKER_COMPOSE_IMPL) $(COMMAND)

.PHONY: docker-compose/build
docker-compose/build:
	$(DOCKER_COMPOSE_IMPL) build

.PHONY: up
up:
	$(DOCKER_COMPOSE_IMPL) up

.PHONY: up-d
up-d:
	$(DOCKER_COMPOSE_IMPL) up -d

.PHONY: up service
up/service:
	$(DOCKER_COMPOSE_IMPL) up $(SERVICE)

.PHONY: bash service
bash/service:
	$(DOCKER_COMPOSE_IMPL) exec $(SERVICE) /bin/bash

.PHONY: down
down:
	$(DOCKER_COMPOSE_IMPL) down

.PHONY: logs
logs:
	$(DOCKER_COMPOSE_IMPL) logs -f

.PHONY: down-remove
__down-remove:
	$(DOCKER_COMPOSE_IMPL) down --rmi all --volumes --remove-orphans
