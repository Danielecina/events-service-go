package repositoryevents

import (
	entities "products-service-go/domains/entities/events"
	"time"

	"github.com/google/uuid"
)

// Create inserts a new event into the PostgreSQL database
func (r *PostgreSQLEventRepository) Create(event *entities.Event) (*entities.Event, error) {
	query := `
		INSERT INTO events (name, description, location, created_at, user_id, event_id)
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id`

	var id int

	eventID := uuid.New().String()
	createdAt := time.Now()

	err := r.db.
		QueryRow(query, event.Name, event.Description, event.Location, createdAt, event.UserID, eventID).
		Scan(&id)

	if err != nil {
		return nil, err
	}

	return event, nil
}
