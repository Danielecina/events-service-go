package middleware

import (
	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
)

// FormatLogger returns a Fiber middleware for logging HTTP requests
func FormatLogger() fiber.Handler {
	return fiberlogger.New(fiberlogger.Config{
		Format: "[APP] [${time}]: ${method} ${path} ${status} - ${latency}\n",
	})
}
