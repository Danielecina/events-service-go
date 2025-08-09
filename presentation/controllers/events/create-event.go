package controllers

import (
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"

	"github.com/gofiber/fiber/v2"
)

// CreateEvent creates a new product event
func (ctrl *EventController) CreateEvent(c *fiber.Ctx) error {
	var eventRequest dto.CreateEventRequest
	if err := c.BodyParser(&eventRequest); err != nil {
		logger.Error("Failed to parse body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateEventErrorResponse{
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    "Failed to parse body",
		})
	}

	logger.Info("Creating new product event with request: %+v", eventRequest)

	event, err := ctrl.createEventsUseCase.Execute(eventRequest)
	if err != nil {
		logger.Error("Failed to create event: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateEventErrorResponse{
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    "Failed to create event",
		})
	}

	logger.Info("Successfully created event with ID: %s", event.EventID)
	return c.Status(fiber.StatusCreated).JSON(event)
}
