.SILENT:

ENV_FILE=.env
include $(ENV_FILE)

DB_SCRIPTS_PATH="scripts"
COMPOSE_FILE_PATH="docker-compose.yml"
TEST_API_DIR="api"
TEST_SERVICE_DIR="internal/app"
DB_MIGRATIONS_PATH="$(DB_SCRIPTS_PATH)"

MYSQL_CONN_STR="mysql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)"
DB_CONN_STR="mysql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

REMOTE_IMAGE_TAG=$(DOCKER_HUB_USERNAME)/$(IMAGE_TAG):0.0.1

ms-up:
	@docker compose --env-file $(ENV_FILE) up -d mysql

ms-down:
	@docker compose --env-file $(ENV_FILE) down -d mysql

pg-up:
	@docker compose --env-file $(ENV_FILE) up -d postgres

pg-down:
	@docker compose --env-file $(ENV_FILE) up -d postgres

pg-downv:
	@docker compose -f $(COMPOSE_FILE_PATH) --env-file $(ENV_FILE) down -v

redis-up:
	@docker compose --env-file $(ENV_FILE) up -d redis

redis-down:
	@docker compose --env-file $(ENV_FILE) down -d redis

migrate-up:
	@migrate -path ./scripts -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?multiStatements=true" up

migrate-down:
	@migrate -path ./scripts -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?multiStatements=true" down

mig-down:
	@migrate -path $(DB_MIGRATIONS_PATH) -database $(DB_CONN_STR) down $c $a

mig-force:
	@migrate -path $(DB_MIGRATIONS_PATH) -database $(DB_CONN_STR) force $v

login:
	@docker login -u $(DOCKER_HUB_USERNAME)

build-local:
	@docker build -t $(IMAGE_TAG) .

build-remote:
	@docker build -t $(REMOTE_IMAGE_TAG) .

push-remote:
	@docker push $(REMOTE_IMAGE_TAG)

pull-remote:
	@docker pull $(REMOTE_IMAGE_TAG)

server-up:
	@docker compose --env-file $(ENV_FILE) up -d server

server-down:
	@docker compose --env-file $(ENV_FILE) down server

server-down-v:
	@docker compose --env-file $(ENV_FILE) down --volumes server

all-up:
	@docker compose --env-file $(ENV_FILE) up -d

all-down:
	@docker compose --env-file $(ENV_FILE) down

all-down-v:
	@docker compose --env-file $(ENV_FILE) down --volumes

test-api:
	@go test ./$(TEST_API_DIR)/...

test-service:
	@go test ./$(TEST_SERVICE_DIR)/...
