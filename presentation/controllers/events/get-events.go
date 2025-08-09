package controllers

import (
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"

	"github.com/gofiber/fiber/v2"
)

// GetEvents returns the list of product events
func (ctrl *EventController) GetEvents(c *fiber.Ctx) error {
	logger.Info("Fetching product events")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	events, err := ctrl.getEventsUseCase.Execute(page, limit)
	if err != nil {
		logger.Error("Failed to fetch events: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.GetEventsErrorResponse{
			StatusCode: fiber.ErrInternalServerError.Code,
			Message:    "Failed to fetch events",
		})
	}

	logger.Info("Successfully fetched %d events", len(events))
	return c.JSON(events)
}
