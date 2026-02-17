package helpers

import (
	"encoding/json"
	"net/http"

	"go-api-template/types"
)

// RespondJSON sends a JSON response
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// RespondSuccess sends a success response
func RespondSuccess(w http.ResponseWriter, message string, data interface{}) {
	RespondJSON(w, http.StatusOK, types.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// RespondError sends an error response
func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, types.APIResponse{
		Success: false,
		Error:   message,
	})
}

// RespondCreated sends a created response
func RespondCreated(w http.ResponseWriter, message string, data interface{}) {
	RespondJSON(w, http.StatusCreated, types.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// RespondPaginated sends a paginated response
func RespondPaginated(w http.ResponseWriter, data interface{}, meta types.Pagination) {
	RespondJSON(w, http.StatusOK, types.PaginatedResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}
