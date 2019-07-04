package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/tylerwray/gus/graphql/fields"
)

// Handler handles all graphql requests coming through the server
var Handler = handler.New(&handler.Config{
	Schema:   &schema,
	Pretty:   true,
	GraphiQL: true,
})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"balance": &fields.Hello,
	},
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createAccessToken": &fields.CreateAccessToken,
	},
})
