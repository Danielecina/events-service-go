package businesscases

import (
	repositoryevents "events-service-go/infrastructure/repositories/events"
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"
)

// GetEventsUseCase handles the business logic for getting events
type GetEventsUseCase struct {
	repo repositoryevents.EventRepositoryClient
}

// Execute retrieves all events
func (uc *GetEventsUseCase) Execute(page int, limit int) ([]dto.GetEventResponse, error) {
	logger.Debug("Starting business case to get Events")

	response, err := uc.repo.GetAll(page, limit)
	if err != nil {
		logger.Error("Failed to get events: %v", err)
		return nil, err
	}

	logger.Info("Successfully retrieved %d events", len(response))
	var events []dto.GetEventResponse
	for _, event := range response {
		events = append(events, dto.GetEventResponse{
			Name:        event.Name,
			Description: event.Description,
			Location:    event.Location,
			UserID:      event.UserID,
			EventID:     event.EventID,
			CreatedAt:   event.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return events, nil
}

// NewGetEventsUseCase creates a new instance of GetEventsUseCase
func NewGetEventsUseCase(repo repositoryevents.EventRepositoryClient) GetEventsUseCase {
	return GetEventsUseCase{
		repo: repo,
	}
}
