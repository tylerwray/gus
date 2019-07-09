package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/tylerwray/gus/api"
)

type resolver struct {
	s *api.Service
}

func newResolvers(s *api.Service) *resolver {
	return &resolver{s}
}

type userBankAccount struct {
	ItemID      string `json:"itemId"`
	AccessToken string `json:"accessToken"`
}

func (r *resolver) linkBankAccount(p graphql.ResolveParams) (interface{}, error) {
	if err := r.s.ValidateAuthToken(p.Context.Value(r.s.AuthTokenKey).(string)); err != nil {
		return nil, err
	}

	publicToken := p.Args["publicToken"]

	res, err := r.s.Plaid.ExchangePublicToken(publicToken.(string))

	if err != nil {
		return nil, err
	}

	return userBankAccount{ItemID: res.ItemID, AccessToken: res.AccessToken}, nil
}

type user struct {
	Username string `json:"username"`
}

func (r *resolver) createUser(p graphql.ResolveParams) (interface{}, error) {
	if err := r.s.ValidateAuthToken(p.Context.Value(r.s.AuthTokenKey).(string)); err != nil {
		return nil, err
	}

	username := p.Args["username"].(string)
	password := p.Args["password"].(string)

	if err := r.s.CreateUser(username, password); err != nil {
		return nil, err
	}

	return user{username}, nil
}
