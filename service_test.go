package main

import (
	"net/http"
	"testing"
)

func TestNewService(t *testing.T) {

	// Call the function.
	_, err := NewService(&http.Client{}, &MockServiceConfig{})
	if err != nil {
		t.Fatalf("Couldn't create service: %v", err)
	}
}
