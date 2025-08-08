package controllers

import (
	"database/sql"
	businesscases "events-service-go/applications/business-cases/events"
	repositoryevents "events-service-go/infrastructure/repositories/events"
	"events-service-go/internal/logger"

	"github.com/gofiber/fiber/v2"
)

// EventController handles HTTP requests for events
type EventController struct {
	getEventsUseCase    *businesscases.GetEventsUseCase
	createEventsUseCase *businesscases.CreateEventsUseCase
}

// NewEventsController factory to create a new events controller
func NewEventsController(
	getEventsUseCase *businesscases.GetEventsUseCase,
	createEventsUseCase *businesscases.CreateEventsUseCase,
) *EventController {
	return &EventController{
		getEventsUseCase:    getEventsUseCase,
		createEventsUseCase: createEventsUseCase,
	}
}

// SetupEventsRoutes configures all event-related routes
func SetupEventsRoutes(app *fiber.App, db *sql.DB) {
	logger.Info("Setting up events routes")

	eventRepo := repositoryevents.NewPostgreSQLEventRepository(db)
	getEventsUseCase := businesscases.NewGetEventsUseCase(eventRepo)
	createEventsUseCase := businesscases.NewCreateEventsUseCase(eventRepo)
	eventsController := NewEventsController(
		getEventsUseCase,
		createEventsUseCase,
	)

	events := app.Group("/events")
	events.Get("/", eventsController.GetEvents)
	events.Post("/", eventsController.CreateEvent)
}
