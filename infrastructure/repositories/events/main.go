package repositoryevents

import (
	"database/sql"
)

// PostgreSQLEventRepository implements the EventRepository interface for PostgreSQL
type PostgreSQLEventRepository struct {
	db *sql.DB
}

// NewPostgreSQLEventRepository creates a new PostgreSQL event repository
func NewPostgreSQLEventRepository(db *sql.DB) *PostgreSQLEventRepository {
	return &PostgreSQLEventRepository{db: db}
}
