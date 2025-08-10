package controllers

import (
	"database/sql"
	businesscases "events-service-go/applications/business-cases/events"
	repositoryevents "events-service-go/infrastructure/repositories/events"
	"events-service-go/internal/logger"

	"github.com/gofiber/fiber/v2"
)

// EventController handles event-related requests
type EventController struct {
	createEventsUseCase businesscases.CreateEventsUseCase
	getEventsUseCase    businesscases.GetEventsUseCase
}

// SetupEventsRoutes configures all event-related routes
func SetupEventsRoutes(app *fiber.App, db *sql.DB) {
	logger.Info("Setting up events routes")

	// Initialize the event repository to pass to the controller
	eventRepo := repositoryevents.NewPostgreSQLEventRepository(db)

	// Create an instance of EventController with the use cases
	eventsController := &EventController{
		createEventsUseCase: businesscases.NewCreateEventsUseCase(eventRepo),
		getEventsUseCase:    businesscases.NewGetEventsUseCase(eventRepo),
	}

	app.Get("/events/", eventsController.GetEvents)
	app.Post("/events/", eventsController.CreateEvent)
}
