package dto

import (
	"encoding/json"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestGetEventsResponse(t *testing.T) {
	t.Run("Expect to serialize GetEventsResponse correctly", func(t *testing.T) {
		resp := []GetEventResponse{
			{
				Name:        "Evento",
				Description: "Descrizione",
				Location:    "Luogo",
				UserID:      "user-1",
			},
		}
		marshalResp, err := json.Marshal(resp)
		require.NoError(t, err)
		snaps.MatchSnapshot(t, string(marshalResp), "get_event_response.json")
	})
}

func TestGetEventsErrorResponse(t *testing.T) {
	t.Run("Expect to serialize GetEventsErrorResponse correctly", func(t *testing.T) {
		errResp := GetEventsErrorResponse{
			Message:    "errore",
			ErrorCode:  "E002",
			StatusCode: 404,
			Errors:     []string{"evento non trovato"},
		}
		marshalInput, err := json.Marshal(errResp)
		require.NoError(t, err)
		snaps.MatchSnapshot(t, string(marshalInput), "get_event_error_response.json")
	})
}
