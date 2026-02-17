package helpers

import (
	"encoding/json"
	"net/http"

	db "go-api-template/database"
)

type ApiFunc func(w http.ResponseWriter, r *http.Request, store db.Storage) error

type APIResponse struct {
	Data    any    `json:"data"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, data any, apiErr *APIError, message string) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)
	}

	response := APIResponse{
		Data:    data,
		Error:   "",
		Message: message,
	}

	if apiErr != nil {
		response.Error = apiErr.Code
		response.Message = apiErr.Message
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to write response"})
	}
}
