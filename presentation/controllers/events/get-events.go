package controllers

import (
	"products-service-go/internal/logger"
	"products-service-go/presentation/dto"

	"github.com/gofiber/fiber/v2"
)

// GetEvents returns the list of product events
func (ctrl *EventController) GetEvents(c *fiber.Ctx) error {
	logger.Info("Fetching product events")

	events, err := ctrl.getEventsUseCase.Execute()
	if err != nil {
		logger.Error("Failed to fetch events: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.GetEventsErrorResponse{
			Message: "Failed to fetch events",
		})
	}

	return c.JSON(events)
}
