package businesscases

import (
	entities "products-service-go/domains/entities/events"
	repositoryevents "products-service-go/infrastructure/repositories/events"
	"products-service-go/presentation/dto"
)

// CreateEventsUseCase handles the business logic for creating events
type CreateEventsUseCase struct {
	eventRepo *repositoryevents.PostgreSQLEventRepository
}

// NewCreateEventsUseCase creates a new create events use case instance
func NewCreateEventsUseCase(eventRepo *repositoryevents.PostgreSQLEventRepository) *CreateEventsUseCase {
	return &CreateEventsUseCase{eventRepo: eventRepo}
}

// Execute creates a new event
func (uc *CreateEventsUseCase) Execute(e *dto.CreateEventRequest) (*entities.Event, error) {
	event, err := uc.eventRepo.Create(&entities.Event{
		Name:        e.Name,
		Description: e.Description,
		Location:    e.Location,
		UserID:      e.UserID,
	})

	if err != nil {
		return nil, err
	}
	return event, nil
}
