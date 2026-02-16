package handlers

import (
	"net/http"
)

// ExampleHandler is a template handler
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from template!"))
}
