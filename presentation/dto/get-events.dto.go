package dto

// GetEventResponse represents the response after getting an event
type GetEventResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location,omitempty"`
	UserID      string `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	EventID     string `json:"event_id"`
}

// GetEventsErrorResponse represents an error response
type GetEventsErrorResponse struct {
	Message    string   `json:"message"`
	ErrorCode  string   `json:"error_code"`
	StatusCode int      `json:"status_code"`
	Errors     []string `json:"errors,omitempty"`
}
