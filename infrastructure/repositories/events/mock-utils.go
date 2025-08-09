package repositoryevents

import entities "events-service-go/domains/entities/events"

// MockEventRepositoryClient Define a mock that implements the EventRepositoryClient interface
type MockEventRepositoryClient struct {
	GetAllFunc func(page int, limit int) ([]entities.Event, error)
	CreateFunc func(event entities.Event) (entities.Event, error)
}

// GetAll MockEventRepositoryClient implements EventRepositoryClient
func (m *MockEventRepositoryClient) GetAll(page int, limit int) ([]entities.Event, error) {
	return m.GetAllFunc(page, limit)
}

// Create MockEventRepositoryClient implements EventRepositoryClient
func (m *MockEventRepositoryClient) Create(event entities.Event) (entities.Event, error) {
	return m.CreateFunc(event)
}
