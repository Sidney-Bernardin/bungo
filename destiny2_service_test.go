package main

import (
	"net/http"
	"testing"
)

func TestGetDestinyManifest(t *testing.T) {

	s, err := NewService(&http.Client{}, &ServiceConfig{ApiKey: getApiKey(t)})
	if err != nil {
		t.Fatalf("couldn't create service: %+v\n", err)
	}

	_, err = s.Destiny2.GetDestinyManifest().Do()
	if err != nil {
		t.Fatalf("couldn't get manifest: %+v\n", err)
	}
}
