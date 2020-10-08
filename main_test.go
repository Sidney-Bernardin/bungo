package gotoorbit

import (
	"os"
	"testing"
)

func getApiKey(t *testing.T) string {

	ret := os.Getenv("API_KEY")
	if ret == "" {
		t.Fatal("Please set the API_KEY environment variable.")
	}

	return ret
}
