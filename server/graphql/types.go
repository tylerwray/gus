package graphql

import "github.com/graphql-go/graphql"

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

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"username": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
