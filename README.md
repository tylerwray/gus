# gus

Backend for the Personal finance management application, Tars

## Getting Started

1. Install go 1.12 `brew install go@1.12`
2. Run `go run .` to start the application

## Migrations

- Migrations will be run when the applications start up.
- We use a simple package called [migrate](https://github.com/golang-migrate/migrate).
- They are written in raw SQL.
- Each migration has a correspoding `<NAME>.up.sql` and `<NAME>.down.sql` file.

If you need to create migrations, install migration cli `brew install golang-migrate`.

Then run:

```
migrate create -ext sql -dir migrations <NAME>
```
