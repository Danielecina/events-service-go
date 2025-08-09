package main

import (
	"events-service-go/infrastructure/databases"
	"events-service-go/internal/logger"
	"events-service-go/internal/middleware"
	controllersevents "events-service-go/presentation/controllers/events"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.Info("Starting application server...")

	db, err := databases.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	app := fiber.New()

	app.Use(middleware.MiddlewareLogger())

	controllersevents.SetupEventsRoutes(app, db)

	logger.Info("Server starting on port 8080...")
	if err := app.Listen(":8080"); err != nil {
		logger.Fatal("Server failed to start: %v", err)
	}
}
