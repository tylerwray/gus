package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/tylerwray/gus/app"
	"github.com/tylerwray/gus/server/handler"
)

// Start the server
func Start() {
	s := app.NewService()

	h := handler.New(s)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	log.Printf("Server is running on %s", addr)

	if err := http.ListenAndServe(addr, cors.Default().Handler(h)); err != nil {
		log.Println("Server could not be started, error:")
		log.Panic(err)
	}
}
