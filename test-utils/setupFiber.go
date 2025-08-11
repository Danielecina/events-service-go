package testutils

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

// SetupFiber initializes a Fiber app with the provided routes and returns it along with a sqlmock instance.
func SetupFiber(
	t *testing.T,
	setupRoutes func(app *fiber.App, db *sql.DB),
) (*fiber.App, sqlmock.Sqlmock) {
	t.Helper()
	db, dbMock, err := sqlmock.New()
	require.NoError(t, err)

	t.Cleanup(func() {
		db.Close()
	})

	app := fiber.New()
	setupRoutes(app, db)
	return app, dbMock
}
