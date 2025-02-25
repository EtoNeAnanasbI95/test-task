CONNECTION_STRING ?= $(PGCONNECT)

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
	@swag i -d ./cmd/ToDoCRUD/,./internal,./models
	@echo "Done"
