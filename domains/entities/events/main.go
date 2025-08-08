package entities

import (
	"time"
)

// Event entity
type Event struct {
	ID          int
	EventID     int
	UserID      int
	Name        string
	Description string
	Location    string
	CreatedAt   time.Time
}

// NewEvent is a Factory function for creating a new event instance
func NewEvent(name, description, location string, createdAt time.Time, userID int) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		UserID:      userID,
	}
}
