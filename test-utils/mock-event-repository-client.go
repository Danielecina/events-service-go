package testutils

import entities "events-service-go/domains/entities/events"

// MockEventRepositoryClient Define a mock that implements the EventRepositoryClient interface
type MockEventRepositoryClient struct {
	GetAllMock func(page int, limit int) ([]entities.Event, error)
	CreateMock func(event entities.Event) (entities.Event, error)
}

// Create MockEventRepositoryClient implements EventRepositoryClient
func (m *MockEventRepositoryClient) Create(event entities.Event) (entities.Event, error) {
	return m.CreateMock(event)
}

// GetAll MockEventRepositoryClient implements EventRepositoryClient
func (m *MockEventRepositoryClient) GetAll(page int, limit int) ([]entities.Event, error) {
	return m.GetAllMock(page, limit)
}
