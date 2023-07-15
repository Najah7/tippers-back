build:
	@docker compose build

up:
	@docker compose up

db:
	@docker compose exec db bash

fmt:
	@gofmt -s -l -w .