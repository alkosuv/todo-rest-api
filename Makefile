user = "postgres"
password = "0000"
host = "192.168.0.20"
port = "5432"
sslmode = disable
db_dev = "todo_dev"
db_test = "todo_test"


db_dev_connect = "postgres://$(user):$(password)@$(host):$(port)/$(db_dev)?sslmode=$(sslmode)"
db_test_connect = "postgres://$(user):$(password)@$(host):$(port)/$(db_test)?sslmode=$(sslmode)"


.PHONY: build
build:
	go build -o bin/server -v cmd/todo/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...


.PHONY: migrate
migrate-create:
	migrate create -ext sql -dir migrations create_todo
migrate-dev-up:
	migrate -path migrations -database $(db_dev_connect) up
migrate-test-up:
	migrate -path migrations -database $(db_test_connect) up
migrate-up:
	make migrate-dev-up
	make migrate-test-up
migrate-dev-down:
	migrate -path migrations -database $(db_dev_connect) down
migrate-test-down:
	migrate -path migrations -database "$(db_test_connect)down
migrate-down:
	make migrate-dev-down
	make migrate-test-down


.PHONY: start
start:
	clear; make build; ./bin/server

.DEFAULT_GOAL := start
