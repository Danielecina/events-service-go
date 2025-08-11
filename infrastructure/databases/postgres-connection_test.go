package databases

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	config := &Config{
		Host:         "host",
		Port:         "1234",
		User:         "user",
		Password:     "pass",
		DBName:       "db",
		SSLMode:      "disable",
		MaxIdleConns: "1",
		MaxOpenConns: "2",
	}
	snaps.MatchSnapshot(t, config.ConnectionString())
}

func TestLoadConfig(t *testing.T) {
	os.Clearenv()
	config := LoadConfig()

	require.Equal(t, "localhost", config.Host)
	require.Equal(t, "5432", config.Port)
	require.Equal(t, "products_user", config.User)
	require.Equal(t, "products_password", config.Password)
	require.Equal(t, "products_service", config.DBName)
	require.Equal(t, "disable", config.SSLMode)
	require.Equal(t, "10", config.MaxIdleConns)
	require.Equal(t, "5", config.MaxOpenConns)
}

func TestConnectDB(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	require.NoError(t, err)

	t.Cleanup(func() {
		db.Close()
	})

	t.Run("Success", func(t *testing.T) {
		sqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
			return db, nil
		}
		dbMock.ExpectPing()
		dbMock.
			ExpectExec("CREATE TABLE IF NOT EXISTS events").
			WillReturnResult(sqlmock.NewResult(1, 1))

		gotDB, err := ConnectDB()
		require.NoError(t, err)
		require.NotNil(t, gotDB)
	})

	t.Run("Connection error", func(t *testing.T) {
		sqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
			return db, nil
		}
		dbMock.ExpectPing()
		dbMock.
			ExpectExec("CREATE TABLE IF NOT EXISTS events").
			WillReturnError(sqlmock.ErrCancelled)

		gotDB, err := ConnectDB()
		require.Error(t, err)
		require.Nil(t, gotDB)
	})
}
