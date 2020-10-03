package main

type ServiceConfigInterface interface {
	ValidateApiKey() error
}

type ServiceConfig struct {
	ApiKey string
}

func (c *ServiceConfig) ValidateApiKey() error {
	return nil
}

// For testing.
type MockServiceConfig struct{}

func (c *MockServiceConfig) ValidateApiKey() error {
	return nil
}
