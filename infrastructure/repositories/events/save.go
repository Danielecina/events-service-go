package repositoryevents

import (
	entities "products-service-go/domains/entities/events"
	"time"
)

// Save creates a new PostgreSQL event repository
func (r *PostgreSQLEventRepository) Save(event *entities.Event) (*entities.Event, error) {
	query := `
        INSERT INTO events (name, description, location, created_at, user_id)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at`

	var id int
	var createdAt time.Time

	err := r.db.QueryRow(query, event.Name, event.Description, event.Location, event.CreatedAt, event.UserID).
		Scan(&id, &createdAt)

	if err != nil {
		return nil, err
	}

	event.ID = id
	return event, nil
}
