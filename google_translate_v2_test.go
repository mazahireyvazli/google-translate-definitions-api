package google_translate_v2_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	google_translate_v2 "github.com/mazahireyvazli/google-translate-definitions-api"
)

func TestParser(t *testing.T) {
	source := google_translate_v2.New()

	entries, err := source.FetchEntries(context.Background(), "benign")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	entriesJSON, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		t.Fatalf("Error marshalling entries to json: %v", err)
	}
	fmt.Println(string(entriesJSON))
}
