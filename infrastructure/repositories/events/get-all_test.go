package repositoryevents

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestEventRepository_GetAll(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		repo := &EventRepository{db: db}
		page := 1
		limit := 2

		rows := sqlmock.
			NewRows([]string{"name", "description", "location", "created_at", "user_id", "event_id"}).
			AddRow("Event1", "Desc1", "Loc1", time.Now(), "user-1", "id-1").
			AddRow("Event2", "Desc2", "Loc2", time.Now(), "user-2", "id-2")

		mock.ExpectQuery(
			regexp.QuoteMeta(`
				SELECT name, description, location, created_at, user_id, event_id 
				FROM events 
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2`,
			),
		).
			WithArgs(limit, (page-1)*limit).
			WillReturnRows(rows)

		events, err := repo.GetAll(page, limit)
		require.NoError(t, err)

		require.Len(t, events, 2)
		require.Equal(t, "Event1", events[0].Name)
		require.Equal(t, "Event2", events[1].Name)

		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("500", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		repo := &EventRepository{db: db}
		page := 1
		limit := 2

		simulatedErr := errors.New("db error")
		mock.ExpectQuery(
			regexp.QuoteMeta(`
				SELECT name, description, location, created_at, user_id, event_id 
        FROM events 
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2`,
			),
		).
			WithArgs(limit, (page-1)*limit).
			WillReturnError(simulatedErr)

		events, err := repo.GetAll(page, limit)
		require.Error(t, err)
		require.Nil(t, events)

		require.NoError(t, mock.ExpectationsWereMet())
	})
}
