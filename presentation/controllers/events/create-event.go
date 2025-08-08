package controllers

import (
	"products-service-go/internal/logger"
	"products-service-go/internal/middleware"
	"products-service-go/presentation/dto"

	"github.com/gofiber/fiber/v2"
)

// CreateEvent creates a new product event
func (ctrl *EventController) CreateEvent(c *fiber.Ctx) error {
	logger.Info("Creating new product event")

	eventRequest, err := middleware.GetValidatedRequest[dto.CreateEventRequest](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateEventErrorResponse{
			Message: err.Error(),
		})
	}

	logger.Info("Starting creating event: %s", eventRequest.Name)
	event, err := ctrl.createEventsUseCase.Execute(&eventRequest)
	if err != nil {
		logger.Error("Failed to create event: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateEventErrorResponse{
			Message: "Failed to create event",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}
