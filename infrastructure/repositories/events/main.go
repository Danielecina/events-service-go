package repositoryevents

import (
	"database/sql"
	entities "events-service-go/domains/entities/events"
	"events-service-go/internal/logger"
)

// EventRepositoryClient defines the methods for event repository
type EventRepositoryClient interface {
	Create(event entities.Event) (entities.Event, error)
	GetAll(page int, limit int) ([]entities.Event, error)
}

// EventRepository implements the EventRepositoryClient interface
type EventRepository struct {
	db *sql.DB
}

// NewPostgreSQLEventRepository creates a new instance of EventRepository
func NewPostgreSQLEventRepository(db *sql.DB) EventRepositoryClient {
	return &EventRepository{db: db}
}

// CreateEventsTable creates the events table in the PostgreSQL database
func CreateEventsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
	  id SERIAL PRIMARY KEY,
		event_id VARCHAR(255) NOT NULL,
		user_id VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		location VARCHAR(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	logger.Debug("Creating events table if not exists")
	_, err := db.Exec(query)
	if err != nil {
		logger.Error("Failed to create events table: %v", err)
		return err
	}

	logger.Info("Events table created successfully")
	return nil
}
