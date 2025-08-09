package businesscases

import (
	repositoryevents "events-service-go/infrastructure/repositories/events"
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"
)

// GetEventsUseCaseClient defines the methods for get events use case
type GetEventsUseCaseClient interface {
	Execute(page int, limit int) ([]dto.GetEventResponse, error)
}

// GetEventsUseCase handles the business logic for getting events
type GetEventsUseCase struct {
	repo repositoryevents.EventRepositoryClient
}

// NewGetEventsUseCase creates a new instance of GetEventsUseCase
func NewGetEventsUseCase(repo repositoryevents.EventRepositoryClient) GetEventsUseCaseClient {
	return &GetEventsUseCase{
		repo: repo,
	}
}

// Execute retrieves all events
func (uc *GetEventsUseCase) Execute(page int, limit int) ([]dto.GetEventResponse, error) {
	logger.Debug("Executing business use case GetEvents on page %d with limit %d", page, limit)

	response, err := uc.repo.GetAll(page, limit)
	if err != nil {
		logger.Error("Failed to get events: %v", err)
		return nil, err
	}

	logger.Info("Successfully retrieved %d events", len(response))
	var dtos []dto.GetEventResponse
	for _, event := range response {
		dtos = append(dtos, dto.GetEventResponse{
			Name:        event.Name,
			Description: event.Description,
			Location:    event.Location,
			UserID:      event.UserID,
		})
	}

	return dtos, nil
}
