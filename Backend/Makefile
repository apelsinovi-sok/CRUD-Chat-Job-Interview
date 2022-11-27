ifeq (migrate,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

define read_env_file
	$(eval ENV_FILE := ./.env)
	@echo " - setup env $(ENV_FILE)"
	$(eval include ./.env)
	$(eval export sed 's/=.*//' ./.env)
endef

# команды для сборки

.PHONY: server
# Запуск сервера нативно. Неоюходимо в .env поменяь хост подключениия к БД на localhost
server:
	go run ./cmd/http/main.go

.PHONY: deploy
# Запуск сервер и базы данных в докере. Работает по умолчанию
deploy:
	docker compose -f .././Docker-Compose.yml -p chat-server up -d

.PHONY: postgres
# Запуск отдельно postgres в докере
postgres:
	docker compose -f ./docker/postgres.yml up -d

.PHONY: migrate
# Запуск миграций goose
#Пример создания миграции таблицы users .././migrate/main.go users sql
# https://github.com/pressly/goose

# $(RUN_ARGS) принимает любые аргументы
migrate:
	go run .././migrate/main.go $(RUN_ARGS)

migrate-up:
	go run .././migrate/main.go up

migrate-down:
	go run .././migrate/main.go down

