package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// GenerateAuthToken creates an auth token to use
func (s *Service) GenerateAuthToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	tokenSecret := os.Getenv("TOKEN_SECRET")

	tokenString, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		return "", errors.New("Could not generate token")
	}

	return tokenString, nil
}

type tokenClaims struct {
	jwt.StandardClaims

	UserID string `json:"user_id"`
}

// ValidateAuthToken validates an auth token
func (s *Service) ValidateAuthToken(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		tokenSecret := os.Getenv("TOKEN_SECRET")

		return []byte(tokenSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := t.Claims.(*tokenClaims); ok && t.Valid {
		return claims.UserID, nil
	}

	return "", nil
}
