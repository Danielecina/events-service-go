package repositoryevents

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestNewPostgreSQLEventRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewPostgreSQLEventRepository(db)
	require.IsType(t, &EventRepository{}, repo)
}

const fullQuery = `
CREATE TABLE IF NOT EXISTS events (
	id SERIAL PRIMARY KEY,
	event_id VARCHAR(255) NOT NULL,
	user_id VARCHAR(255) NOT NULL,
	name VARCHAR(255) NOT NULL,
	description TEXT,
	location VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

func TestCreateEventsTable(t *testing.T) {
	t.Run("Expect to create table successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(fullQuery)).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err = CreateEventsTable(db)
		require.NoError(t, err)

		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Expect to fail creation of table", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		mock.
			ExpectExec(regexp.QuoteMeta(fullQuery)).
			WillReturnError(errors.New("db error"))

		err = CreateEventsTable(db)
		require.Error(t, err)

		require.NoError(t, mock.ExpectationsWereMet())
	})
}
