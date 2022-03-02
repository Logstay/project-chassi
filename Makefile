.PHONY: migrate migrate.down run tests

OS_NAME := $(shell uname)

ifeq ($(OS_NAME), Darwin)
	DB_HOST := "docker.for.mac.localhost"
else
	DB_HOST := "localhost"
endif

migrate:
	@docker container run --rm --name migrate -v ${PWD}/migrations:/migrations --network host migrate/migrate \
    -path=/migrations \
		-database postgres://postgres:postgres00@$(DB_HOST):5435/escola_biblica?sslmode=disable up

migrate.down:
	@docker container run --rm --name migrate -v ${PWD}/migrations:/migrations --network host migrate/migrate \
    -path=/migrations \
		-database postgres://postgres:postgres00@$(DB_HOST):5435/escola_biblica?sslmode=disable down

run:
	@go run cmd/main.go

tests:
	@go test -v ./...
