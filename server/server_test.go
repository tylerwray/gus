package server_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/tylerwray/gus/app"
	"github.com/tylerwray/gus/server"
)

func TestNewHandler(t *testing.T) {
	s := app.NewService()
	handler := server.NewHandler(s)
	ts := httptest.NewServer(handler)
	defer ts.Close()

	fmt.Println(ts.URL)
}
