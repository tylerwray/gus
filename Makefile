include .env

db-reset: db-drop
	@migrate --source file://migrations --database ${DATABASE_URL} up


db-drop:
	@migrate --source file://migrations --database ${DATABASE_URL} drop

generate-migration: 
	@migrate create --ext sql --dir migrations ${NAME}
