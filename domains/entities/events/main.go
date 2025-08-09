package entities

import (
	"time"
)

// Event entity
type Event struct {
	EventID     string
	UserID      string
	Name        string
	Description string
	Location    string
	CreatedAt   time.Time
}

// NewEvent is a Factory function for creating a new event instance
func NewEvent(name, description, location string, createdAt time.Time, userID string) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		UserID:      userID,
	}
}
