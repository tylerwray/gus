package app

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // Postgres driver
	"github.com/plaid/plaid-go/plaid"
)

// Service ...
type Service struct {
	db    *sql.DB
	plaid *plaid.Client
}

// NewService creates a new service instance
func NewService() *Service {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Panic(err)
	}

	// Check that our connection is good
	if err := db.Ping(); err != nil {
		log.Panic(err)
	}

	clientOptions := plaid.ClientOptions{
		ClientID:    os.Getenv("PLAID_CLIENT_ID"),
		Secret:      os.Getenv("PLAID_SECRET"),
		PublicKey:   os.Getenv("PLAID_PUBLIC_KEY"),
		Environment: plaid.Development, // Available environments are Sandbox, Development, and Production
	}

	plaidClient, err := plaid.NewClient(clientOptions)

	if err != nil {
		log.Panic("Couldn't create plaid client")
	}

	return &Service{
		db,
		plaidClient,
	}
}
