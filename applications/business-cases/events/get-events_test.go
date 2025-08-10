package businesscases

import (
	"fmt"
	"testing"

	entities "events-service-go/domains/entities/events"
	repositoryevents "events-service-go/infrastructure/repositories/events"

	"github.com/stretchr/testify/require"
)

func TestGetAllEventsUseCase(t *testing.T) {
	t.Run("Execute_Success", func(t *testing.T) {
		expectedEvents := []entities.Event{
			{
				EventID:     "event-1",
				Name:        "Event 1",
				Description: "Desc 1",
				Location:    "Loc 1",
				UserID:      "user-1",
			},
			{
				EventID:     "event-2",
				Name:        "Event 2",
				Description: "Desc 2",
				Location:    "Loc 2",
				UserID:      "user-2",
			},
		}

		eventsRepositories := &repositoryevents.MockEventRepositoryClient{
			GetAllMock: func(page int, limit int) ([]entities.Event, error) {
				return expectedEvents, nil
			},
		}

		events, err := eventsRepositories.GetAll(1, 10)

		require.NoError(t, err)
		require.Equal(t, expectedEvents, events)
	})

	t.Run("Execute_Error", func(t *testing.T) {
		eventsRepositories := &repositoryevents.MockEventRepositoryClient{
			GetAllMock: func(page int, limit int) ([]entities.Event, error) {
				return nil, fmt.Errorf("db error")
			},
		}

		_, err := eventsRepositories.GetAll(1, 10)
		require.Error(t, err)
	})
}
