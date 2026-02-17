package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ParseJSON parses JSON request body
func ParseJSON(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		return errors.New("request body is empty")
	}

	return json.Unmarshal(body, v)
}

// ValidateRequired checks if a string field is not empty
func ValidateRequired(field, name string) error {
	if field == "" {
		return errors.New(name + " is required")
	}
	return nil
}

// ValidateEmail checks if email is valid (basic check)
func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	// TODO: Add proper email validation
	// You can use a library like govalidator or implement regex
	return nil
}

// ValidateMinLength checks minimum length
func ValidateMinLength(field, name string, min int) error {
	if len(field) < min {
		return errors.New(name + " must be at least " + string(rune(min)) + " characters")
	}
	return nil
}
