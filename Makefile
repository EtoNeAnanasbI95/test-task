.PHONY: gen-mock
.gen-mock:
	mockery --name=ClientInterface --dir=api --output=internal/service/mocks --case=underscore

.PHONY: migrations-up
migrations-up:
	# "postgres://user:password@localhost:5432/dbname?sslmode=disable"
	migrate -database "$(CONNECTION_STRING)" -path migrations up

.PHONY: migrations-down
migrations-down:
	migrate -database "$(CONNECTION_STRING)" -path migrations down

.PHONY: gen-swag
gen-swag:
	@echo "Generate swagger docs"
	@swag init -g cmd/test-task/main.go
	@echo "Done"

.PHONY: api-migrate
api-migrate:
	@echo "Applying migrations"
	@go run ./cmd/migrate/main.go --config "$(CONFIG_PATH)" --migrations "$(MIGRATIONS_PATH)"
	@echo "Done"

.PHONY: api-run
api-run: api-migrate
	@echo "Starting API"
	@go run ./cmd/test-task/main.go --config "$(CONFIG_PATH)"
	@echo "Done"
