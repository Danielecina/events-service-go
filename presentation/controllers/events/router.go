package controllers

import (
	"database/sql"
	businesscases "products-service-go/applications/business-cases/events"
	repositoryevents "products-service-go/infrastructure/repositories/events"
	"products-service-go/internal/logger"
	"products-service-go/internal/middleware"
	"products-service-go/presentation/dto"

	"github.com/gofiber/fiber/v2"
)

// EventController handles HTTP requests for events
type EventController struct {
	getEventsUseCase    *businesscases.GetEventsUseCase
	createEventsUseCase *businesscases.CreateEventsUseCase
}

// NewEventsController factory to create a new events controller
func NewEventsController(
	getEventsUseCase *businesscases.GetEventsUseCase,
	createEventsUseCase *businesscases.CreateEventsUseCase,
) *EventController {
	return &EventController{
		getEventsUseCase:    getEventsUseCase,
		createEventsUseCase: createEventsUseCase,
	}
}

// SetupEventsRoutes configures all event-related routes
func SetupEventsRoutes(app fiber.Router, db *sql.DB) {
	eventRepo := repositoryevents.NewPostgreSQLEventRepository(db)
	getEventsUseCase := businesscases.NewGetEventsUseCase(eventRepo)
	createEventsUseCase := businesscases.NewCreateEventsUseCase(eventRepo)
	eventsController := NewEventsController(
		getEventsUseCase,
		createEventsUseCase,
	)

	events := app.Group("/events")
	logger.Info("Setting up events routes")

	events.Get("/",
		middleware.ValidateWithSchema(middleware.SchemaConfig{
			ResponseDTOs: map[int]interface{}{
				200: dto.GetEventsResponse{},
				500: dto.GetEventsErrorResponse{},
			},
		}),
		eventsController.GetEvents,
	)

	events.Post("/",
		middleware.ValidateWithSchema(middleware.SchemaConfig{
			RequestDTO: dto.CreateEventRequest{},
			ResponseDTOs: map[int]interface{}{
				201: dto.CreateEventResponse{},
				500: dto.CreateEventErrorResponse{},
			},
		}),
		eventsController.CreateEvent,
	)
}
