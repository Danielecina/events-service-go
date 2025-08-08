package controllers

import (
	"products-service-go/internal/logger"

	"github.com/gofiber/fiber/v2"
)

// SetupHealthRoutes configures all health check routes
func SetupHealthRoutes(app fiber.Router) {
	health := app.Group("/health")

	logger.Info("Setting up health check routes")
	health.Get("/readiness", func(c *fiber.Ctx) error {
		logger.Info("Readiness check requested from IP: %s", c.IP())
		return c.SendStatus(fiber.StatusOK)
	})

	health.Get("/liveness", func(c *fiber.Ctx) error {
		logger.Info("Liveness check requested from IP: %s", c.IP())
		return c.SendStatus(fiber.StatusOK)
	})
}
