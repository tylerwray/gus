package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/tylerwray/gus/plaid"
)

type exchange struct {
	ItemID      string `json:"itemId"`
	AccessToken string `json:"accessToken"`
}

var exchangeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Exchange",
		Fields: graphql.Fields{
			"itemId": &graphql.Field{
				Type: graphql.String,
			},
			"accessToken": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// CreateAccessToken field in the schema
var CreateAccessToken = graphql.Field{
	Type: exchangeType,
	Args: graphql.FieldConfigArgument{
		"publicToken": &graphql.ArgumentConfig{
			Description: "The public token given back from link",
			Type:        graphql.String,
		},
	},
	Description: "Exchange a plaid public token for an access token and item id",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
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
	},
}
