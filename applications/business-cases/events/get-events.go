package businesscases

import (
	repositoryevents "events-service-go/infrastructure/repositories/events"
	"events-service-go/internal/logger"
	"events-service-go/presentation/dto"
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
func (uc *GetEventsUseCase) Execute(page int, limit int) ([]dto.GetEventResponse, error) {
	logger.Debug("Executing business use case GetEvents on page %d with limit %d", page, limit)

	response, err := uc.eventRepo.GetAll(page, limit)
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
