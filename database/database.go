package database

import (
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Migration driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // Migration file source
	_ "github.com/lib/pq"                                      // postgres driver
	"golang.org/x/crypto/bcrypt"
)

var connStr = ""

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Migrate the database
func Migrate() error {
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

// GenerateHash a salted hash for the input string
func GenerateHash(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

// CompareHash a hash and a string
func CompareHash(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
