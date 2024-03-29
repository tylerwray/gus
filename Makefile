include .env

db.migrate:
	@migrate --source file://migrations --database ${DATABASE_URL} up

db.reset: db.drop db.migrate

db.drop:
	@migrate --source file://migrations --database ${DATABASE_URL} drop

db.seed:
	psql ${DATABASE_URL} -a -f migrations/seeds.sql

db.create:
	psql postgres://postgres:postgres@localhost:5432?sslmode=disable -c 'CREATE DATABASE gus_dev'

gen.migration: 
	@migrate create --ext sql --dir migrations ${NAME}

test:
	go test ./...