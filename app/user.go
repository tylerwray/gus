package app

import "strings"

// CreateUser service creates a user
func (s *Service) CreateUser(username, password string) error {
	hash, err := generateHash(password)

	if err != nil {
		return err
	}

	if _, err := s.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", strings.ToLower(username), hash); err != nil {
		return err
	}

	return nil
}

// User type
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUser returns a user based on their username
func (s *Service) GetUser(username string) (User, error) {
	var hash string

	if err := s.db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hash); err != nil {
		return User{}, err
	}

	return User{Username: username, Password: hash}, nil
}
