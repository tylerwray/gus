package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/tylerwray/gus/app"
)

func authTokenMiddleware(s *app.Service, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), " ")[1]

		userID, err := s.ValidateAuthToken(token)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"Invalid token"}`))
			return
		}

		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userIDContextKey, userID)))
	})
}
