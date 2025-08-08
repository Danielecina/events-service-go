package dto

import "time"

// CreateEventRequest represents the request body for creating an event
type CreateEventRequest struct {
	ID          int       `json:"id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"required,min=10,max=500"`
	CreatedAt   time.Time `json:"created_at" validate:"required"`
	Location    string    `json:"location,omitempty" validate:"omitempty"`
	UserID      int       `json:"user_id" validate:"required,min=1"`
}

// CreateEventResponse represents the response after creating an event
type CreateEventResponse struct {
	ID int `json:"event_id,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message    string   `json:"message"`
	ErrorCode  string   `json:"error_code"`
	StatusCode int      `json:"status_code"`
	Errors     []string `json:"errors,omitempty"`
}
