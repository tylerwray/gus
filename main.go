package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/tylerwray/gus/auth"
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

	mux.Handle("/api/v1/login", auth.LoginHandler(db))
	mux.Handle("/graphql", auth.TokenMiddleware(gqlHandler))

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	log.Printf("Server is running on %s", addr)
	http.ListenAndServe(addr, cors.Default().Handler(mux))
}
