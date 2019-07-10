package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// GenerateAuthToken creates an auth token to use
func (s *Service) GenerateAuthToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	tokenSecret := os.Getenv("TOKEN_SECRET")

	tokenString, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		return "", errors.New("Could not generate token")
	}

	return tokenString, nil
}

// ValidateAuthToken validates an auth token
func (s *Service) ValidateAuthToken(token string) error {
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
