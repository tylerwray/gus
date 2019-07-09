package main

import (
	"os"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Migration driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // Migration file source
)

func migrateDB() error {
	m, err := migrate.New(
		"file://migrations",
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		return err
	}

	if err := m.Up(); err != migrate.ErrNoChange {
		return err
	}

	return nil
}
