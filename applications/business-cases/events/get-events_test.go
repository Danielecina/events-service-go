package businesscases

import (
	"fmt"
	"testing"

	entities "events-service-go/domains/entities/events"
	testutils "events-service-go/test-utils"

	"github.com/gkampitakis/go-snaps/snaps"
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

		eventsRepositories := &testutils.MockEventRepositoryClient{
			GetAllMock: func(page int, limit int) ([]entities.Event, error) {
				return expectedEvents, nil
			},
		}

		useCase := NewGetEventsUseCase(eventsRepositories)

		response, err := useCase.Execute(1, 10)
		require.NoError(t, err)
		require.Len(t, response, 2)
		snaps.MatchSnapshot(t, response)
	})

	t.Run("Execute_Error", func(t *testing.T) {
		eventsRepositories := &testutils.MockEventRepositoryClient{
			GetAllMock: func(page int, limit int) ([]entities.Event, error) {
				return nil, fmt.Errorf("db error")
			},
		}

		useCase := NewGetEventsUseCase(eventsRepositories)
		got, err := useCase.Execute(1, 10)
		require.Error(t, err)
		require.Nil(t, got)
	})
}
