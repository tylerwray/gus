package graphql

import (
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/tylerwray/gus/app"
)

// NewHandler creates a handler to handle all graphql requests
func NewHandler(s *app.Service) http.Handler {
	schema := newSchema(s)

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}
