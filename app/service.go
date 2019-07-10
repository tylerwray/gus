package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/plaid/plaid-go/plaid"
)

type authTokenKey struct{}

// Service ...
type Service struct {
	db           *sql.DB
	Plaid        *plaid.Client
	AuthTokenKey *authTokenKey
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
		&authTokenKey{},
	}
}
