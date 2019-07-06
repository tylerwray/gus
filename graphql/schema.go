package graphql

import (
	"database/sql"

	"github.com/graphql-go/graphql"
)

// NewSchema creates a schema graphql object
func NewSchema(db *sql.DB) (graphql.Schema, error) {
	resolvers := newResolvers(db)

	var queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "world", nil
				},
			},
		},
	})

	createAccessToken := graphql.Field{
		Type: exchangeType,
		Args: graphql.FieldConfigArgument{
			"publicToken": &graphql.ArgumentConfig{
				Description: "The public token given back from link",
				Type:        graphql.String,
			},
		},
		Description: "Exchange a plaid public token for an access token and item id",
		Resolve:     resolvers.createAccessToken,
	}

	createUser := graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Description: "The username of the user.",
				Type:        graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Description: "The password of the user.",
				Type:        graphql.String,
			},
		},
		Description: "Create a user.",
		Resolve:     resolvers.createUser,
	}

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createAccessToken": &createAccessToken,
			"createUser":        &createUser,
		},
	})

	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)
}
