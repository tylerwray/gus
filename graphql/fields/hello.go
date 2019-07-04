package fields

import "github.com/graphql-go/graphql"

// Hello field in the schema
var Hello = graphql.Field{
	Type:        graphql.String,
	Description: "Hello, World!",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "world!", nil
	},
}
