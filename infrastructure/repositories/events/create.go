package repositoryevents

import (
	entities "events-service-go/domains/entities/events"
	"events-service-go/internal/logger"
	"time"

	"github.com/google/uuid"
)

// Create inserts a new event into the PostgreSQL database
func (r *PostgreSQLEventRepository) Create(event entities.Event) (entities.Event, error) {
	logger.Debug("Executing repository method Create for event: %s", event.Name)

	query := `
	   INSERT INTO events (name, description, location, created_at, user_id, event_id)
	   VALUES ($1, $2, $3, $4, $5, $6)`

	eventID := uuid.New().String()
	createdAt := time.Now()
	event.EventID = eventID
	event.CreatedAt = createdAt

	_, err := r.db.Exec(query, event.Name, event.Description, event.Location, createdAt, event.UserID, eventID)
	if err != nil {
		logger.Error("Failed to execute query: %v", err)
		return entities.Event{}, err
	}

	return event, nil
}
