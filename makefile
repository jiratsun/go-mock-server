SHELL := /bin/bash

schema:
	@echo "Generating schema with Go Jet"
	jet -dsn=mysql://root:rootpassword@/go_mock_server -path=./.jet-gen -ignore-tables=schema_migrations

up:
	@echo "Migrating up"
	migrate -path=./migrations -database=mysql://root:rootpassword@/go_mock_server up

down:
	@echo "Migrating down"
	migrate -path=./migrations -database=mysql://root:rootpassword@/go_mock_server down

run:
	@echo "Starting Go server"
	go run /Users/jiratviriyataranon/Desktop/my-projects/go-mock-server/src local
