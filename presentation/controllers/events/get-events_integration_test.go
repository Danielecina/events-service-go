package controllers

import (
	"encoding/json"
	"events-service-go/presentation/dto"
	testutils "events-service-go/presentation/test-utils"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

const getEventsQuery = `
SELECT name, description, location, created_at, user_id, event_id 
FROM events 
ORDER BY created_at DESC
LIMIT $1 OFFSET $2`

func TestEventController(t *testing.T) {
	app, dbMock := testutils.SetupFiber(t, SetupEventsRoutes)

	t.Run("200", func(t *testing.T) {
		page := 1
		limit := 2
		offset := (page - 1) * limit

		dbMock.
			ExpectQuery(regexp.QuoteMeta(getEventsQuery)).
			WithArgs(limit, offset).
			WillReturnRows(sqlmock.
				NewRows([]string{"name", "description", "location", "created_at", "user_id", "event_id"}).
				AddRow("Event 1", "Description 1", "Location 1", testutils.TimeParser("2025-08-11T10:00:00Z"), "user-1", "event-1").
				AddRow("Event 2", "Description 2", "Location 2", testutils.TimeParser("2025-08-11T11:00:00Z"), "user-2", "event-2"),
			)

		req, err := http.NewRequest("GET", "/events", nil)
		require.NoError(t, err)
		q := req.URL.Query()
		q.Add("page", fmt.Sprintf("%d", page))
		q.Add("limit", fmt.Sprintf("%d", limit))
		req.URL.RawQuery = q.Encode()

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusOK, resp.StatusCode)

		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var out []dto.GetEventResponse
		err = json.Unmarshal(bodyBytes, &out)
		require.NoError(t, err)
		require.Equal(t, []dto.GetEventResponse{
			{
				Name:        "Event 1",
				Description: "Description 1",
				Location:    "Location 1",
				UserID:      "user-1",
				CreatedAt:   "2025-08-11T10:00:00Z",
				EventID:     "event-1",
			},
			{
				Name:        "Event 2",
				Description: "Description 2",
				Location:    "Location 2",
				UserID:      "user-2",
				CreatedAt:   "2025-08-11T11:00:00Z",
				EventID:     "event-2",
			},
		}, out)
	})

	t.Run("500", func(t *testing.T) {
		dbMock.
			ExpectQuery(regexp.QuoteMeta(getEventsQuery)).
			WithArgs(1, 10).
			WillReturnError(sqlmock.ErrCancelled)

		req, err := http.NewRequest("GET", "/events", nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var out dto.GetEventsErrorResponse
		err = json.Unmarshal(bodyBytes, &out)
		require.NoError(t, err)
		require.Equal(t, dto.GetEventsErrorResponse{
			StatusCode: 500,
			Message:    "Failed to fetch events",
		}, out)
	})
}
