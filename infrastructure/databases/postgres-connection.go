package databases

import (
	"database/sql"
	"events-service-go/internal/utils"
	"fmt"
	"log"
	"strconv"

	repositoryevents "events-service-go/infrastructure/repositories/events"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

// Config holds the configuration for the PostgreSQL connection.
type Config struct {
	Host         string
	Port         string
	User         string
	Password     string
	DBName       string
	SSLMode      string
	MaxIdleConns string
	MaxOpenConns string
}

// LoadConfig loads the PostgreSQL configuration from environment variables or defaults.
func LoadConfig() *Config {
	return &Config{
		Host:         utils.GetEnv("DB_HOST", "localhost"),
		Port:         utils.GetEnv("DB_PORT", "5432"),
		User:         utils.GetEnv("DB_USER", "products_user"),
		Password:     utils.GetEnv("DB_PASSWORD", "products_password"),
		DBName:       utils.GetEnv("DB_NAME", "products_service"),
		SSLMode:      utils.GetEnv("DB_SSLMODE", "disable"),
		MaxIdleConns: utils.GetEnv("DB_MAX_IDLE_CONNS", "10"),
		MaxOpenConns: utils.GetEnv("DB_MAX_OPEN_CONNS", "5"),
	}
}

// ConnectionString returns the connection string for the database.
func (c *Config) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

// ConnectDB establishes a connection to the PostgreSQL database.
func ConnectDB() (*sql.DB, error) {
	config := LoadConfig()
	db, err := sql.Open("postgres", config.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := repositoryevents.CreateEventsTable(db); err != nil {
		return nil, fmt.Errorf("failed to create events table: %w", err)
	}

	maxIdleConns, err := strconv.Atoi(config.MaxIdleConns)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_MAX_IDLE_CONNS value: %w", err)
	}

	maxOpenConns, err := strconv.Atoi(config.MaxOpenConns)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_MAX_OPEN_CONNS value: %w", err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	log.Println("Successfully connected to PostgreSQL")
	return db, nil
}
