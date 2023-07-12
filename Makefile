.PHONY: migrate migrate.down run tests dc_up dc_down

migrate:
	@docker container run --rm --name migrate -v ${PWD}/migrations:/migrations --network host migrate/migrate \
    -path=/migrations \
		-database postgres://postgres:postgres00@localhost:5435/escola_biblica?sslmode=disable up

migrate.down:
	@docker container run --rm --name migrate -v ${PWD}/migrations:/migrations --network host migrate/migrate \
    -path=/migrations \
		-database postgres://postgres:postgres00@localhost:5435/escola_biblica?sslmode=disable down

run:
	@go run cmd/main.go

tests:
	@go test -v ./...

dc_up:
	@docker compose up -d

dc_down:
	@docker compose down