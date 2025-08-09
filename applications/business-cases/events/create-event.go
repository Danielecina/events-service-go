package businesscases

import (
	entities "events-service-go/domains/entities/events"
	repositoryevents "events-service-go/infrastructure/repositories/events"
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"
)

// CreateEventsUseCaseClient defines the methods for create events use case
type CreateEventsUseCaseClient interface {
	Execute(e dto.CreateEventRequest) (dto.CreateEventResponse, error)
}

// CreateEventsUseCase handles the business logic for creating events
type CreateEventsUseCase struct {
	repo repositoryevents.EventRepositoryClient
}

// NewCreateEventsUseCase creates a new instance of CreateEventsUseCase
func NewCreateEventsUseCase(repo repositoryevents.EventRepositoryClient) CreateEventsUseCaseClient {
	return &CreateEventsUseCase{
		repo: repo,
	}
}

// Execute creates a new event
func (uc *CreateEventsUseCase) Execute(e dto.CreateEventRequest) (dto.CreateEventResponse, error) {
	logger.Debug("Executing business use case CreateEvent with %+v", e)
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
