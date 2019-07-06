package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/tylerwray/gus/database"
	"github.com/tylerwray/gus/graphql"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db, err := database.New()

	if err != nil {
		log.Fatal(err)
	}

	if err = database.Migrate(); err != nil {
		log.Fatal(err)
	}

	schema, err := graphql.NewSchema(db)

	if err != nil {
		log.Fatal(err)
	}

	gqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	mux := http.NewServeMux()
	mux.Handle("/graphql", gqlHandler)

	log.Print("Server is running on port 8080")
	http.ListenAndServe(":8080", cors.Default().Handler(mux))
}
