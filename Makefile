CURRENT_DIR := $(shell pwd)

swag:
	swag init -g cmd/main.go --parseDependency --parseInternal

# migrate:
# 	go run ./migrate/migrate.go

run:
	 go run cmd/main.go
