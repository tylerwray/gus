package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tylerwray/gus/app"
	"github.com/tylerwray/gus/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	if err := migrateDB(); err != nil {
		log.Fatal(err)
	}

	s := app.NewService()

	handler := server.NewHandler(s)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	log.Printf("Server is running on %s", addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Println("Server could not be started, error:")
		log.Panic(err)
	}
}
