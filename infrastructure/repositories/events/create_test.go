package repositoryevents

import (
	"errors"
	"regexp"
	"testing"
	"time"

	entities "events-service-go/domains/entities/events"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestCreateEventRepository(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		repo := &EventRepository{db: db}
		event := entities.Event{
			Name:        "Test Event",
			Description: "Test Description",
			Location:    "Test Location",
			UserID:      "user-123",
		}

		mock.ExpectExec(
			regexp.QuoteMeta(`
				INSERT INTO events (name, description, location, created_at, user_id, event_id)
				VALUES ($1, $2, $3, $4, $5, $6)`,
			),
		).
			WithArgs(event.Name, event.Description, event.Location, sqlmock.AnyArg(), event.UserID, sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		createdEvent, err := repo.Create(event)
		require.NoError(t, err)

		require.Equal(t, event.Name, createdEvent.Name)
		require.Equal(t, event.Description, createdEvent.Description)
		require.Equal(t, event.Location, createdEvent.Location)
		require.Equal(t, event.UserID, createdEvent.UserID)
		require.NotEmpty(t, createdEvent.EventID)
		require.WithinDuration(t, time.Now(), createdEvent.CreatedAt, time.Second)

		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("500", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		repo := &EventRepository{db: db}
		event := entities.Event{
			Name:        "Test Event",
			Description: "Test Description",
			Location:    "Test Location",
			UserID:      "user-123",
		}

		simulatedErr := errors.New("db error")
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO events (name, description, location, created_at, user_id, event_id)
               VALUES ($1, $2, $3, $4, $5, $6)`)).
			WithArgs(event.Name, event.Description, event.Location, sqlmock.AnyArg(), event.UserID, sqlmock.AnyArg()).
			WillReturnError(simulatedErr)

		createdEvent, err := repo.Create(event)
		require.Error(t, err)
		require.Equal(t, entities.Event{}, createdEvent)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
