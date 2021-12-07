MIGRATE_DB_DSN=postgres://postgres:postgres@localhost/migrate_playground?sslmode=disable
all: help

init:
	# create new user
	# create new database migrate_playground 
## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## migrate/up: migrate up
migrate/up:
	migrate -path db/migration -database ${MIGRATE_DB_DSN} -verbose up

## migrate/down: migrate down
migrate/down:
	migrate -path db/migration -database ${MIGRATE_DB_DSN} -verbose down

.PHONY: help init migrate/up migrate/down
