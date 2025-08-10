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
