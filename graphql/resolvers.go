package graphql

import (
	"database/sql"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/tylerwray/gus/auth"
	"github.com/tylerwray/gus/database"
	"github.com/tylerwray/gus/plaid"
)

type resolver struct {
	db *sql.DB
}

func newResolvers(db *sql.DB) *resolver {
	return &resolver{db}
}

type exchange struct {
	ItemID      string `json:"itemId"`
	AccessToken string `json:"accessToken"`
}

func (r *resolver) createAccessToken(p graphql.ResolveParams) (interface{}, error) {
	if err := auth.ValidateToken(p.Context.Value(auth.TokenKey).(string)); err != nil {
		return nil, err
	}

	publicToken := p.Args["publicToken"]

	client, err := plaid.New()

	if err != nil {
		return nil, err
	}

	res, err := client.ExchangePublicToken(publicToken.(string))

	if err != nil {
		return nil, err
	}

	return exchange{ItemID: res.ItemID, AccessToken: res.AccessToken}, nil
}

type user struct {
	Username string `json:"username"`
}

func (r *resolver) createUser(p graphql.ResolveParams) (interface{}, error) {
	if err := auth.ValidateToken(p.Context.Value(auth.TokenKey).(string)); err != nil {
		return nil, err
	}

	// TODO: Extract this into a `CreateUser` service
	username := strings.ToLower(p.Args["username"].(string))
	password, err := database.GenerateHash(p.Args["password"].(string))

	if err != nil {
		return nil, err
	}

	if _, err := r.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password); err != nil {
		return nil, err
	}

	return user{username}, nil
}
