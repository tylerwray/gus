package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tylerwray/gus/api"
)

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login authorizes a user and sends them a token
func Login(s *api.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)

		var body loginBody

		if err := decoder.Decode(&body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "JSON body could not be read"}`))
			return
		}

		user, err := s.GetUser(body.Username)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Incorrect login information"}`))
			return
		}

		if err := s.ComparePasswordHash(user.Password, body.Password); err != nil {
			w.WriteHeader(401)
			w.Write([]byte(`{"error": "Incorrect login information"}`))
			return
		}

		token, err := s.GenerateAuthToken()

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(`{"error": "Could not generate token"}`))
			return
		}

		w.Write([]byte(fmt.Sprintf(`{"data":{"token": "%s"}}`, token)))
	}
}
