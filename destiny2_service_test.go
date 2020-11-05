package bungo

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"testing"
)

var (
	apiKey = flag.String("apikey", "", "Your bungie api key.")
	lookAt = flag.String("lookAt", "", "Set lookAt to the name of a test function to see the response. Example: go test -lookAt 'TestGetDestinyManifest' .")
	pretty = flag.Bool("pretty", false, "If you are using lookAt, you can set pretty to format the response.")
)

func TestGetDestinyManifest(t *testing.T) {

	// Create service.
	s, err := NewService(&http.Client{})
	if err != nil {
		t.Fatalf("couln't create service: %v", err)
	}

	// Get destiny manifest.
	res, err := s.Destiny2.GetDestinyManifest().Do()
	if err != nil {
		t.Fatalf("couldn't get manifest: %v", err)
	}

	// Print response.
	if *lookAt == "TestGetDestinyManifest" {

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
