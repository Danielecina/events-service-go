package controllers

import (
	"products-service-go/internal/logger"

	"github.com/gofiber/fiber/v2"
)

// SetupEventsRoutes configures all event-related routes
func SetupEventsRoutes(app fiber.Router, ctrl *EventController) {
	events := app.Group("/events")
	logger.Info("Setting up events routes")

	events.Get("/", ctrl.GetEvents)
	// events.Post("/", middleware.ValidateSchema[entities.Event](), CreateEvent)
}
