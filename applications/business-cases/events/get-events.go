package businesscases

import (
	entities "events-service-go/domains/entities/events"
	repositoryevents "events-service-go/infrastructure/repositories/events"
)

// GetEventsUseCase handles the business logic for getting events
type GetEventsUseCase struct {
	eventRepo *repositoryevents.PostgreSQLEventRepository
}

// NewGetEventsUseCase creates a new get events use case instance
func NewGetEventsUseCase(eventRepo *repositoryevents.PostgreSQLEventRepository) *GetEventsUseCase {
	return &GetEventsUseCase{eventRepo: eventRepo}
}

// Execute retrieves all events
func (uc *GetEventsUseCase) Execute() ([]entities.Event, error) {
	return uc.eventRepo.GetAll()
}
