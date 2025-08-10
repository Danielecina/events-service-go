package businesscases

import (
	"encoding/json"
	"fmt"
	"testing"

	entities "events-service-go/domains/entities/events"
	repositoryevents "events-service-go/infrastructure/repositories/events"
	"events-service-go/presentation/dto"

	"github.com/stretchr/testify/require"
)

func TestCreateEventsUseCase(t *testing.T) {
	t.Run("Execute_Success", func(t *testing.T) {
		eventID := "test-event-id"
		eventsRepositories := &repositoryevents.MockEventRepositoryClient{
			CreateMock: func(event entities.Event) (entities.Event, error) {
				return entities.Event{
					EventID:     eventID,
					Name:        event.Name,
					Description: event.Description,
					Location:    event.Location,
					UserID:      event.UserID,
				}, nil
			},
		}

		useCase := NewCreateEventUseCase(eventsRepositories)

		input := dto.CreateEventRequest{
			Name:        "Test Event",
			Description: "Test Desc",
			Location:    "Test Loc",
			UserID:      "user-1",
		}

		resp, err := useCase.Execute(input)

		fmt.Println("Response:", resp)

		require.NoError(t, err)
		respBytes, _ := json.Marshal(resp)
		require.JSONEq(t, string(respBytes), `{"event_id": "test-event-id"}`)
	})

	t.Run("Execute_Error", func(t *testing.T) {
		eventsRepositories := &repositoryevents.MockEventRepositoryClient{
			CreateMock: func(event entities.Event) (entities.Event, error) {
				return entities.Event{}, fmt.Errorf("db error")
			},
		}

		useCase := NewCreateEventUseCase(eventsRepositories)

		input := dto.CreateEventRequest{
			Name:        "Test Event",
			Description: "Test Desc",
			Location:    "Test Loc",
			UserID:      "user-1",
		}

		_, err := useCase.Execute(input)
		require.Error(t, err)
	})
}
