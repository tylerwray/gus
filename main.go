package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/tylerwray/gus/api"
	"github.com/tylerwray/gus/graphql"
	"github.com/tylerwray/gus/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	if err := migrateDB(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	s := api.NewService()

	gqlHandler := graphql.NewHandler(s)

	mux.HandleFunc("/api/v1/login", handlers.Login(s))
	mux.Handle("/graphql", authTokenMiddleware(s, gqlHandler))

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	log.Printf("Server is running on %s", addr)

	http.ListenAndServe(addr, cors.Default().Handler(mux))
}

func authTokenMiddleware(s *api.Service, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), " ")[1]

		if err := s.ValidateAuthToken(token); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"Invalid token"}`))
			return
		}

		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), s.AuthTokenKey, token)))
	})
}
