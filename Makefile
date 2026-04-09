include .env
export $(shell sed 's/=.*//' .env)

MIGRATIONS_PATH = ./cmd/migrate/migrations 

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up 
migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_ADDR)" up

.PHONY: migrate-down
migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_ADDR)" down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go

.PHONY: migrate-force
migrate-force:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_ADDR)" force $(filter-out $@,$(MAKECMDGOALS))
