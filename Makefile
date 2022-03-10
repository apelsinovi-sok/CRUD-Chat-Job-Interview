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
# Запуск локального сервера
server:
	go run ./cmd/http/main.go

.PHONY: deploy
# Запуск приложения в докере (не работает) (также необходимо изменить строку подключения к базе) (пример в редми)
deploy:
	docker-compose -f ./my-network.yml up


.PHONY: postgres
# Запуск отдельно postgres в докере
postgres:
	docker-compose -f ./docker/postgres.yml up -d

.PHONY: migrate
# Запуск миграций goose
# https://github.com/pressly/goose
migrate:
	go run ./cmd/migrate/main.go $(RUN_ARGS)

migrate-up:
	go run ./cmd/migrate/main.go up

migrate-down:
	go run ./cmd/migrate/main.go down

