package controllers

import (
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"

	"github.com/gofiber/fiber/v2"
)

// CreateEvent creates a new product event
func (ctrl *EventController) CreateEvent(c *fiber.Ctx) error {
	logger.Info("Creating new product event")

	var eventRequest dto.CreateEventRequest
	if err := c.BodyParser(&eventRequest); err != nil {
		logger.Error("Failed to parse body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateEventErrorResponse{
			StatusCode: fiber.ErrBadRequest.Code,
			Message:    "Failed to parse body",
		})
	}

	event, err := ctrl.createEventsUseCase.Execute(eventRequest)
	if err != nil {
		logger.Error("Failed to create event: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateEventErrorResponse{
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    "Failed to create event",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}
