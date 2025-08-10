package dto

import (
	"encoding/json"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestCreateEventRequest(t *testing.T) {
	t.Run("Expect to serialize CreateEventRequest correctly", func(t *testing.T) {
		input := CreateEventRequest{
			Name:        "Evento",
			Description: "Descrizione",
			Location:    "Luogo",
			UserID:      "user-1",
		}
		marshalInput, err := json.Marshal(input)
		require.NoError(t, err)
		expectedStr := string(marshalInput)
		snaps.MatchSnapshot(t, expectedStr, "create_event_request.json")
	})
}

func TestCreateEventErrorResponse(t *testing.T) {
	t.Run("Expect to serialize CreateEventErrorResponse correctly", func(t *testing.T) {
		errResp := CreateEventErrorResponse{
			Message:    "errore",
			ErrorCode:  "E001",
			StatusCode: 400,
			Errors:     []string{"campo mancante"},
		}
		marshalInput, err := json.Marshal(errResp)
		require.NoError(t, err)
		expectedStr := string(marshalInput)
		snaps.MatchSnapshot(t, expectedStr, "create_event_error_response.json")
	})
}
