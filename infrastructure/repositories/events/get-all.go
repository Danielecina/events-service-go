package repositoryevents

import (
	entities "events-service-go/domains/entities/events"
	"events-service-go/internal/logger"
)

// GetAll retrieves all events from the PostgreSQL database
func (r *PostgreSQLEventRepository) GetAll(page int, limit int) ([]entities.Event, error) {
	logger.Debug("Executing repository method GetAll with page %d and limit %d", page, limit)

	query := `
		SELECT name, description, location, created_at, user_id, event_id 
		FROM events 
		ORDER BY created_at 
		DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, (page-1)*limit)

	if err != nil {
		logger.Error("Failed to execute query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var events []entities.Event

	for rows.Next() {
		var event entities.Event
		err := rows.Scan(&event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UserID, &event.EventID)

		if err != nil {
			logger.Error("Failed to scan row: %v", err)
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
