package google_translate_v2_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	google_translate_v2 "github.com/mazahireyvazli/google-translate-definitions-api"
)

func TestParser(t *testing.T) {
	source := google_translate_v2.New(google_translate_v2.WithRequestOptions(&google_translate_v2.RequestOptions{
		From:   google_translate_v2.English,
		To:     google_translate_v2.Azerbaijani,
		HL:     google_translate_v2.English,
		TLD:    "com",
		RPCIDs: "MkEWBc",
	}))

	entries, err := source.FetchEntries(context.Background(), "primitive")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	entriesJSON, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		t.Fatalf("Error marshalling entries to json: %v", err)
	}
	fmt.Println(string(entriesJSON))
}

func TestSentence01(t *testing.T) {
	source := google_translate_v2.New(google_translate_v2.WithRequestOptions(&google_translate_v2.RequestOptions{
		From:   google_translate_v2.English,
		To:     google_translate_v2.German,
		HL:     google_translate_v2.English,
		TLD:    "com",
		RPCIDs: "MkEWBc",
	}))

	entries, err := source.FetchEntries(context.Background(), "I will have gone by the time you get there")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	entriesJSON, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		t.Fatalf("Error marshalling entries to json: %v", err)
	}
	fmt.Println(string(entriesJSON))
}

func TestIdiom(t *testing.T) {
	source := google_translate_v2.New(google_translate_v2.WithRequestOptions(&google_translate_v2.RequestOptions{
		From:   google_translate_v2.English,
		To:     google_translate_v2.Turkish,
		HL:     google_translate_v2.English,
		TLD:    "com",
		RPCIDs: "MkEWBc",
	}))

	entries, err := source.FetchEntries(context.Background(), "Come off it.")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	entriesJSON, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		t.Fatalf("Error marshalling entries to json: %v", err)
	}
	fmt.Println(string(entriesJSON))
}
