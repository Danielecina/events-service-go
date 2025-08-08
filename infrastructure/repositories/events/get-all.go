package repositoryevents

import entities "events-service-go/domains/entities/events"

// GetAll retrieves all events from the PostgreSQL database
func (r *PostgreSQLEventRepository) GetAll() ([]entities.Event, error) {
	query := `
		SELECT id, name, description, location, date_time, user_id 
		FROM events 
		ORDER BY date_time 
		DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []entities.Event

	for rows.Next() {
		var event entities.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
