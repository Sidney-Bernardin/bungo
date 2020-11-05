package bungo

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"testing"
)

var (
	apiKey = flag.String("apiKey", "", "Your bungie api key.")
	lookAt = flag.String("lookAt", "", "Set lookAt to the name of a test function to see the response. Example: go test -lookAt 'TestGetDestinyManifest' .")
	pretty = flag.Bool("pretty", false, "If you are using lookAt, you can set pretty to format the response.")
)

func printResponse(t *testing.T, res interface{}, funcName string) {

	if *lookAt == funcName {

		if *pretty {

			// Marshal indent.
			marshaled, err := json.MarshalIndent(res, "", "  ")
			if err != nil {
				t.Fatalf("couldn't marshal-indent response: %s", err)
			}

			fmt.Printf("%s\n", string(marshaled))
			return
		}

		// Marshal
		marshaled, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("couldn't marshal response: %s", err)
		}
		fmt.Printf("%s\n", string(marshaled))
	}
}

func TestGetDestinyManifest(t *testing.T) {

	// Create service.
	s, err := NewService(&http.Client{}, *apiKey)
	if err != nil {
		t.Fatalf("couln't create service: %v", err)
	}

	// Get destiny manifest.
	res, err := s.Destiny2.GetDestinyManifest().Do()
	if err != nil {
		t.Fatalf("couldn't get manifest: %v", err)
	}

	// Print response.
	printResponse(t, res, "TestGetDestinyManifest")
}

func TestGetDestinyEntityDefinition(t *testing.T) {

	// Create service.
	s, err := NewService(&http.Client{}, *apiKey)
	if err != nil {
		t.Fatalf("couln't create service: %v", err)
	}

	// Get destiny entity definition.
	res, err := s.Destiny2.GetDestinyEntityDefinition(
		"DestinyEquipmentSlotDefinition",
		"3448274439", // helmet bucket hash
	).Do()
	if err != nil {
		t.Fatalf("couldn't get entity definition: %v", err)
	}

	// Print response.
	printResponse(t, res, "TestGetDestinyEntityDefinition")
}

func TestSearchDestinyPlayer(t *testing.T) {

	tables := []struct {
		returnOriginalProfile string
	}{
		{"true"},
		{"false"},
	}

	// Create service.
	s, err := NewService(&http.Client{}, *apiKey)
	if err != nil {
		t.Fatalf("couln't create service: %v", err)
	}

	// Run test cases.
	for _, table := range tables {
		res, err := s.Destiny2.SearchDestinyPlayer("2", "a_neutrino").
			ReturnOriginalProfile(table.returnOriginalProfile).
			Do()

		if err != nil {
			t.Fatalf("couldn't get entity definition: %v", err)
		}

		// Print response.
		printResponse(t, res, "TestSearchDestinyPlayer")
	}
}

func TestGetLinkedProfiles(t *testing.T) {

	tables := []struct {
		getAllMemberships string
	}{
		{"true"},
		{"false"},
	}

	// Create service.
	s, err := NewService(&http.Client{}, *apiKey)
	if err != nil {
		t.Fatalf("couln't create service: %v", err)
	}

	// Run test cases.
	for _, table := range tables {
		res, err := s.Destiny2.GetLinkedProfiles("2", "4611686018458149036").
			GetAllMemeberships(table.getAllMemberships).
			Do()

		if err != nil {
			t.Fatalf("couldn't get linked profiles: %v", err)
		}

		// Print response.
		printResponse(t, res, "TestSearchDestinyPlayer")
	}
}

func TestGetCharacter(t *testing.T) {

	tables := []struct {
		components string
	}{
		{"200"},
		{"Characters"},
		{"205"},
		{"CharacterEquipment"},
	}

	// Create service.
	s, err := NewService(&http.Client{}, *apiKey)
	if err != nil {
		t.Fatalf("couln't create service: %v", err)
	}

	// Run test cases.
	for _, table := range tables {
		res, err := s.Destiny2.GetCharacter(
			"2",
			"4611686018458149036",
			"2305843009294908257").
			Components(table.components).
			Do()

		if err != nil {
			t.Fatalf("couldn't get character: %v\ncomponents: %s", err, table.components)
		}

		// Print response.
		printResponse(t, res, "TestGetCharacter")
	}
}
