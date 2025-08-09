package businesscases

import (
	entities "events-service-go/domains/entities/events"
	repositoryevents "events-service-go/infrastructure/repositories/events"
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"
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
func (u *CreateEventsUseCase) Execute(e dto.CreateEventRequest) (dto.CreateEventResponse, error) {
	logger.Debug("Executing business use case CreateEvent with %+v", e)
	event, err := u.eventRepo.Create(entities.Event{
		Name:        e.Name,
		Description: e.Description,
		Location:    e.Location,
		UserID:      e.UserID,
	})

	if err != nil {
		return dto.CreateEventResponse{}, err
	}

	return dto.CreateEventResponse{
		EventID: event.EventID,
	}, nil
}
