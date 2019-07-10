package handler

import (
	"net/http"

	"github.com/tylerwray/gus/app"
	"github.com/tylerwray/gus/server/graphql"
)

// New creates a the root handler
func New(s *app.Service) http.Handler {
	gqlHandler := graphql.NewHandler(s)
	handler := http.NewServeMux()

	handler.HandleFunc("/api/v1/login", login(s))
	handler.Handle("/graphql", authTokenMiddleware(s, gqlHandler))

	return handler
}
