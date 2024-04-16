CASSANDRA_MIGRATIONS = ./migrations/cassandra
POSTGRES_MIGRATIONS = ./migrations/postgres
MONGO_MIGRATIONS = ./migrations/mongo

include .env
export

# CASSANDRA


.PHONY: start-cassandra
start-cassandra:
	(cd ./docker-composes/cassandra; docker-compose up -d --build)
	make migrate-up-cassandra

.PHONY: stop-cassandra
stop-cassandra:
	(cd ./docker-composes/cassandra; docker-compose stop)

.PHONY: migrate-up-cassandra
migrate-up-cassandra:
	migrate -path=$(CASSANDRA_MIGRATIONS) -database "${CASSANDRA_MIGRATE_ENDPOINT}" up

.PHONY: migrate-down-cassandra
migrate-down-cassandra:
	migrate -path=$(CASSANDRA_MIGRATIONS) -database "${CASSANDRA_MIGRATE_ENDPOINT}" down

.PHONY: migrate-new-cassandra
migrate-new-cassandra:
	migrate create -ext cql -dir migrations/cassandra -seq story



# POSTGRES

.PHONY: start-postgres
start-postgres:
	(cd ./docker-composes/postgres; docker-compose up -d --build)
	make migrate-up-postgres

.PHONY: stop-postgres
stop-postgres:
	(cd ./docker-composes/postgres; docker-compose stop)

.PHONY: migrate-down-postgres
migrate-down-postgres:
	migrate -path=$(POSTGRES_MIGRATIONS) -database "${POSTGRES_MIGRATE_ENDPOINT}" down

.PHONY: migrate-up-postgres
migrate-up-postgres:
	migrate -path=$(POSTGRES_MIGRATIONS) -database "${POSTGRES_MIGRATE_ENDPOINT}" up

.PHONY: migrate-new-postgres
migrate-new-postgres:
	migrate create -ext sql -dir migrations/postgres -seq story_views_stat


# MONGO


.PHONY: start-mongo
start-mongo:
	(cd ./docker-composes/mongo; docker-compose up -d --build)
	make migrate-up-mongo

.PHONY: stop-mongo
stop-mongo:
	(cd ./docker-composes/mongo; docker-compose stop)

.PHONY: migrate-up-mongo
migrate-up-mongo:
	migrate -path=$(MONGO_MIGRATIONS) -database "${MONGO_MIGRATE_ENDPOINT}" up

.PHONY: migrate-new-mongo
migrate-new-mongo:
	migrate create -ext json -dir migrations/mongo -seq story_stat

lint:
	golangci-lint run ./...