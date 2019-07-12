package server

import (
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/tylerwray/gus/app"
)

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(s *app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate HTTP Method
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Must POST to /login."}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// Parse body
		decoder := json.NewDecoder(r.Body)

		var body loginBody

		if err := decoder.Decode(&body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "JSON body could not be read"}`))
			return
		}

		// Get user
		user, err := s.GetUser(body.Username)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Incorrect login information"}`))
			return
		}

		// Check request password against user's password
		if err := s.ComparePasswordHash(user.Password, body.Password); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Incorrect login information"}`))
			return
		}

		// Respond with token
		token, err := s.GenerateAuthToken(user.ID)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Could not generate token"}`))
			return
		}

		w.Write([]byte(fmt.Sprintf(`{"data":{"token": "%s"}}`, token)))
	}
}

type signUpBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signUp(s *app.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate HTTP Method
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Must POST to /sign-up."}`))
			return
		}

		// Set response content-type
		w.Header().Set("Content-Type", "application/json")

		// Parse body
		decoder := json.NewDecoder(r.Body)

		var body signUpBody

		if err := decoder.Decode(&body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "JSON body could not be read."}`))
			return
		}

		// Create User
		err := s.CreateUser(body.Username, body.Password)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Failed to create user."}`))
			return
		}

		// Respond with Token
		userID := getUserID(r.Context())

		fmt.Println(userID)

		token, err := s.GenerateAuthToken(userID)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Could not generate token"}`))
			return
		}

		w.Write([]byte(fmt.Sprintf(`{"data":{"token": "%s"}}`, token)))
	}
}
