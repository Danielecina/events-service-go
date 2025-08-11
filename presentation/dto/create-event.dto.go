package dto

// CreateEventRequest represents the request body for creating an event
type CreateEventRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location,omitempty"`
	UserID      string `json:"user_id"`
}

// CreateEventResponse represents the response after creating an event
type CreateEventResponse struct {
	EventID string `json:"event_id"`
}

// CreateEventErrorResponse represents an error response
type CreateEventErrorResponse struct {
	Message    string   `json:"message"`
	ErrorCode  string   `json:"error_code"`
	StatusCode int      `json:"status_code"`
	Errors     []string `json:"errors,omitempty"`
}
