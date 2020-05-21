.PHONY: build
build:
	go build -o bin/server -v cmd/todo/main.go


.PHONY: migrate
migrate-create:
	migrate create -ext sql -dir migrations create_todo
migrate-dev-up:
	migrate -path migrations -database "postgres://postgres:0601@192.168.0.10:5432/todo_dev?sslmode=disable" up
migrate-test-up:
	migrate -path migrations -database "postgres://postgres:0601@192.168.0.10:5432/todo_test?sslmode=disable" up
migrate-up:
	make migrate-dev-up
	make migrate-test-up
migrate-dev-down:
	migrate -path migrations -database "postgres://postgres:0601@192.168.0.10:5432/todo_dev?sslmode=disable" down
migrate-test-down:
	migrate -path migrations -database "postgres://postgres:0601@192.168.0.10:5432/todo_test?sslmode=disable" down
migrate-down:
	make migrate-dev-down
	make migrate-test-down


.PHONY: start
start:
	clear; make build; ./bin/server

.DEFAULT_GOAL := start
