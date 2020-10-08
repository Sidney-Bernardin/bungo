package bungo

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

func TestGetEntityDefinition(t *testing.T) {

	s, err := NewService(&http.Client{}, &ServiceConfig{ApiKey: getApiKey(t)})
	if err != nil {
		t.Fatalf("couldn't create service: %+v\n", err)
	}

	_, err = s.Destiny2.GetDestinyEntityDefinition("DestinyInventoryItemDefinition", "720351794").Do()

	if err != nil {
		t.Fatalf("couldn't get item: %+v\n", err)
	}
}

func TestSearchDestinyPlayer(t *testing.T) {

	s, err := NewService(&http.Client{}, &ServiceConfig{ApiKey: getApiKey(t)})
	if err != nil {
		t.Fatalf("couldn't create service: %+v\n", err)
	}

	_, err = s.Destiny2.SearchDestinyPlayer("2", "a_neutrino").Do()
	if err != nil {
		t.Fatalf("couldn't get destiny player: %+v\n", err)
	}

	_, err = s.Destiny2.SearchDestinyPlayer("2", "a_neutrino").
		ReturnOriginalProfile("true").
		Do()

	if err != nil {
		t.Fatalf("couldn't get destiny player's OG profile: %+v\n", err)
	}
}
