package plaid

import (
	"os"

	"github.com/plaid/plaid-go/plaid"
)

// New plaid client
func New() (*plaid.Client, error) {
	clientOptions := plaid.ClientOptions{
		ClientID:    os.Getenv("PLAID_CLIENT_ID"),
		Secret:      os.Getenv("PLAID_SECRET"),
		PublicKey:   os.Getenv("PLAID_PUBLIC_KEY"),
		Environment: plaid.Development, // Available environments are Sandbox, Development, and Production
	}

	client, err := plaid.NewClient(clientOptions)

	if err != nil {
		return nil, err
	}

	return client, nil
}
