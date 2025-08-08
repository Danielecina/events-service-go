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

// NewEvent creates a new event instance
func NewEvent(name, description, location string, createdAt time.Time, userID int) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		CreatedAt:   createdAt,
		UserID:      userID,
	}
}

func (e *Event) IsValid() bool {
	return e.Name != "" && e.UserID > 0
}

func (e *Event) IsInFuture() bool {
	return e.CreatedAt.After(time.Now())
}
