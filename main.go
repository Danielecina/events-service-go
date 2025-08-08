package main

import (
	"log"
	"products-service-go/infrastructure/databases"
	"products-service-go/internal/logger"
	"products-service-go/internal/middleware"
	controllersevents "products-service-go/presentation/controllers/events"
	controllershealthz "products-service-go/presentation/controllers/healthz"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.Info("Starting application server...")

	db, err := databases.ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	app := fiber.New()

	app.Use(middleware.MiddlewareLogger())

	controllershealthz.SetupHealthRoutes(app)
	controllersevents.SetupEventsRoutes(app, db)

	logger.Info("Server starting on port 8080...")
	if err := app.Listen(":8080"); err != nil {
		logger.Fatal("Server failed to start: %v", err)
	}
}
