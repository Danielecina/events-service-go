package controllers

import (
	businesscases "products-service-go/applications/business-cases/events"
	"products-service-go/internal/logger"
	"products-service-go/presentation/dto"

	"github.com/gofiber/fiber/v2"
)

// EventController handles HTTP requests for events
type EventController struct {
	getEventsUseCase *businesscases.GetEventsUseCase
}

// NewEventController creates a new event controller
func NewEventController(getEventsUseCase *businesscases.GetEventsUseCase) *EventController {
	return &EventController{
		getEventsUseCase: getEventsUseCase,
	}
}

// GetEvents returns the list of product events
func (ctrl *EventController) GetEvents(c *fiber.Ctx) error {
	logger.Info("Fetching product events")

	events, err := ctrl.getEventsUseCase.Execute()
	if err != nil {
		logger.Error("Failed to fetch events: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Failed to fetch events",
		})
	}

	return c.JSON(events)
}
