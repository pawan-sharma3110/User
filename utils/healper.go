package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"user/model"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// ParseJson parses JSON from the request body
func ParseJson(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// ValidateUser performs basic validation on a user
func ValidateUser(user model.User) error {
	if user.Email == "" || user.Password == "" {
		return fmt.Errorf("email and password cannot be empty")
	}
	return nil
}

// WriteError writes an error message to the response
func WriteError(w http.ResponseWriter, status int, err error) {
	http.Error(w, err.Error(), status)
}

// WriteJson writes a JSON response
func WriteJson(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
