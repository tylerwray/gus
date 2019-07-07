package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/tylerwray/gus/database"
)

// Generate a valid token
func Generate() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	tokenSecret := os.Getenv("TOKEN_SECRET")

	tokenString, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		return "", errors.New("Could not generate token")
	}

	return tokenString, nil
}

type tokenKey struct{}

// TokenKey is the key for the request context to access the token
var TokenKey = &tokenKey{}

// ValidateToken validates an auth token
func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		tokenSecret := os.Getenv("TOKEN_SECRET")

		return []byte(tokenSecret), nil
	})

	return err
}

// TokenMiddleware is a middleware that puts the token on the request context
func TokenMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), " ")[1]

		if err := ValidateToken(token); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"Invalid token"}`))
			return
		}

		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), TokenKey, token)))
	})
}

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler authorizes a user and sends them a token
func LoginHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)

		var body loginBody

		if err := decoder.Decode(&body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "JSON body could not be read"}`))
			return
		}

		var hash string

		err := db.QueryRow("SELECT password FROM users WHERE username = $1", body.Username).Scan(&hash)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Incorrect login information"}`))
			return
		}

		if err := database.CompareHash(hash, body.Password); err != nil {
			w.WriteHeader(401)
			w.Write([]byte(`{"error": "Incorrect login information"}`))
			return
		}

		token, err := Generate()

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(`{"error": "Could not generate token"}`))
			return
		}

		w.Write([]byte(fmt.Sprintf(`{"data":{"token": "%s"}}`, token)))
	})
}
