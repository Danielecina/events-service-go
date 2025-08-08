package dto

import entities "products-service-go/domains/entities/events"

// GetEventsResponse represents the response after getting events
type GetEventsResponse struct {
	Events []entities.Event `json:"events"`
	ID     int              `json:"id"`
}

// GetEventsErrorResponse represents an error response
type GetEventsErrorResponse struct {
	Message    string   `json:"message"`
	ErrorCode  string   `json:"error_code"`
	StatusCode int      `json:"status_code"`
	Errors     []string `json:"errors,omitempty"`
}
