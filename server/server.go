package server

import (
	"context"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/rs/cors"
	"github.com/tylerwray/gus/app"
)

type userIDContextKeyType struct{}

var userIDContextKey *userIDContextKeyType

func getUserID(ctx context.Context) string {
	return ctx.Value(userIDContextKey).(string)
}

// NewHandler creates a new http.Handler for the server
func NewHandler(s *app.Service) http.Handler {
	userIDContextKey = &userIDContextKeyType{}

	schema := newSchema(s)

	gqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	handler := http.NewServeMux()

	handler.HandleFunc("/api/v1/login", login(s))
	handler.HandleFunc("/api/v1/sign-up", signUp(s))
	handler.Handle("/graphql", authTokenMiddleware(s, gqlHandler))

	return cors.Default().Handler(handler)
}
