package businesscases

import (
	entities "events-service-go/domains/entities/events"
	repositoryevents "events-service-go/infrastructure/repositories/events"
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"
)

// CreateEventsUseCase handles the business logic for creating events
type CreateEventsUseCase struct {
	repo repositoryevents.EventRepositoryClient
}

// Execute creates a new event
func (uc *CreateEventsUseCase) Execute(e dto.CreateEventRequest) (dto.CreateEventResponse, error) {
	logger.Debug("Starting business case to create Event")
	event, err := uc.repo.Create(entities.Event{
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

// NewCreateEventUseCase creates a new instance of CreateEventsUseCase
func NewCreateEventUseCase(repo repositoryevents.EventRepositoryClient) CreateEventsUseCase {
	return CreateEventsUseCase{
		repo: repo,
	}
}
