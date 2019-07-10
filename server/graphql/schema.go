package graphql

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/tylerwray/gus/app"
)

func newSchema(s *app.Service) graphql.Schema {
	resolvers := newResolvers(s)

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

	var linkBankAccount = graphql.Field{
		Type: exchangeType,
		Args: graphql.FieldConfigArgument{
			"publicToken": &graphql.ArgumentConfig{
				Description: "The public token given back from link",
				Type:        graphql.String,
			},
		},
		Description: "Link a bank account by exchanging a plaid public token for an access token and item id",
		Resolve:     resolvers.linkBankAccount,
	}

	var createUser = graphql.Field{
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
			"linkBankAccount": &linkBankAccount,
			"createUser":      &createUser,
		},
	})

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)

	if err != nil {
		log.Panic(err)
	}

	return schema
}
