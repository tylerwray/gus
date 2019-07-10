package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/tylerwray/gus/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	if err := migrateDB(); err != nil {
		log.Fatal(err)
	}

	server.Start()
}