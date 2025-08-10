package dto

// GetEventsResponse represents the response after getting events
type GetEventsResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location,omitempty"`
	UserID      string `json:"user_id"`
}

// GetEventsErrorResponse represents an error response
type GetEventsErrorResponse struct {
	Message    string   `json:"message"`
	ErrorCode  string   `json:"error_code"`
	StatusCode int      `json:"status_code"`
	Errors     []string `json:"errors,omitempty"`
}
