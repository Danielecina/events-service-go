package controllers

import (
	"encoding/json"
	"events-service-go/presentation/dto"
	testutils "events-service-go/test-utils"
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestEventController_CreateEvent(t *testing.T) {
	app, dbMock := testutils.SetupFiber(t, SetupEventsRoutes)

	t.Run("201", func(t *testing.T) {
		defer gock.Off()

		reqBody := dto.CreateEventRequest{
			Name:        "test",
			Description: "desc",
			Location:    "loc",
			UserID:      "user-1",
		}

		dbMock.ExpectExec(
			regexp.QuoteMeta(`
				INSERT INTO events (name, description, location, created_at, user_id, event_id)
				VALUES ($1, $2, $3, $4, $5, $6)`,
			),
		).
			WithArgs(reqBody.Name, reqBody.Description, reqBody.Location, sqlmock.AnyArg(), reqBody.UserID, sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		jsonBody, err := json.Marshal(reqBody)
		req, err := http.NewRequest("POST", "/events", strings.NewReader(string(jsonBody)))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		require.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})

	t.Run("400 - with invalid request body", func(t *testing.T) {
		notValidBodyReq := strings.NewReader("not-json")
		req, err := http.NewRequest("POST", "/events", notValidBodyReq)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("500 - with database error", func(t *testing.T) {
		defer gock.Off()

		reqBody := dto.CreateEventRequest{
			Name:        "test",
			Description: "desc",
			Location:    "loc",
			UserID:      "user-1",
		}

		dbMock.ExpectExec(
			regexp.QuoteMeta(`
				INSERT INTO events (name, description, location, created_at, user_id, event_id)
				VALUES ($1, $2, $3, $4, $5, $6)`,
			),
		).
			WithArgs(reqBody.Name, reqBody.Description, reqBody.Location, sqlmock.AnyArg(), reqBody.UserID, sqlmock.AnyArg()).
			WillReturnError(sqlmock.ErrCancelled)

		jsonBody, err := json.Marshal(reqBody)
		req, err := http.NewRequest("POST", "/events", strings.NewReader(string(jsonBody)))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var errResp dto.CreateEventErrorResponse
		err = json.Unmarshal(bodyBytes, &errResp)
		require.NoError(t, err)

		require.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		require.Equal(t, dto.CreateEventErrorResponse{
			StatusCode: 500,
			Message:    "Failed to create event",
		}, errResp)
	})
}
