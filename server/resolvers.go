package server

import (
	"github.com/graphql-go/graphql"
	"github.com/tylerwray/gus/app"
)

type resolver struct {
	s *app.Service
}

func newResolvers(s *app.Service) *resolver {
	return &resolver{s}
}

type userBankAccount struct {
	ItemID      string `json:"itemId"`
	AccessToken string `json:"accessToken"`
}

func (r *resolver) linkBankAccount(p graphql.ResolveParams) (interface{}, error) {
	var userID = "12"
	publicToken := p.Args["publicToken"]

	err := r.s.AddBankAccount(userID, publicToken.(string))

	if err != nil {
		return false, err
	}

	return true, nil
}

type user struct {
	Username string `json:"username"`
}

func (r *resolver) createUser(p graphql.ResolveParams) (interface{}, error) {
	username := p.Args["username"].(string)
	password := p.Args["password"].(string)

	if err := r.s.CreateUser(username, password); err != nil {
		return nil, err
	}

	return user{username}, nil
}
