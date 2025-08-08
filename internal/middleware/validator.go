package middleware

import (
	"products-service-go/internal/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateSchema creates a middleware that automatically validates request body against struct validate tags
func ValidateSchema[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger.Info("Validation middleware started")
		var data T

		// Parse JSON body into struct
		if err := c.BodyParser(&data); err != nil {
			logger.Error("Failed to parse JSON body: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid JSON format",
			})
		}

		// Validate struct using validate tags
		if err := validate.Struct(&data); err != nil {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.Error())
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Validation failed",
				"errors":  errors,
			})
		}

		logger.Info("Validation passed, storing data in context")
		// Store validated data in context for handler to use
		c.Locals("validatedData", &data)

		return c.Next()
	}
}

// GetValidatedData retrieves the validated data from context
func GetValidatedData[T any](c *fiber.Ctx) T {
	data := c.Locals("validatedData")
	if data == nil {
		var zero T
		return zero
	}
	return data.(T)
}
