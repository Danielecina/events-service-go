package repositoryevents

import (
	"database/sql"
	entities "events-service-go/domains/entities/events"
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
