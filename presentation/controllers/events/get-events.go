package controllers

import (
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"

	"github.com/gofiber/fiber/v2"
)

// GetEvents returns the list of product events
func (ctrl *EventController) GetEvents(c *fiber.Ctx) error {
	logger.Info("Fetching product events")

	events, err := ctrl.getEventsUseCase.Execute()
	if err != nil {
		logger.Error("Failed to fetch events: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.GetEventsErrorResponse{
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    "Failed to fetch events",
		})
	}

	return c.JSON(events)
}
