package types

// APIResponse is a generic API response structure
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginatedResponse for paginated results
type PaginatedResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Meta    Pagination  `json:"meta"`
}

// Pagination metadata
type Pagination struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}
