package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/tylerwray/gus/graphql"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/graphql", graphql.Handler)

	handler := cors.Default().Handler(mux)

	log.Print("Server is running on port 8080")
	http.ListenAndServe(":8080", handler)
}
