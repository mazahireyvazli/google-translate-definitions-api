package google_translate_v2_test

import (
	"context"
	"testing"

	google_translate_v2 "github.com/mazahireyvazli/google-translate-definitions-api"
)

func TestParser(t *testing.T) {
	source := google_translate_v2.New(google_translate_v2.WithTranslationOptions(&google_translate_v2.TranslationOptions{
		From: google_translate_v2.English,
		To:   google_translate_v2.Azerbaijani,
		HL:   google_translate_v2.English,
		TLD:  "com",
	}))

	entries, err := source.FetchEntries(context.Background(), "primitive")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	if len(entries) == 0 {
		t.Fatalf("No entries found")
	}

	for _, entry := range entries {
		if entry.Text != "primitive" {
			t.Fatalf("Expected entry text to be 'primitive', but got '%s'", entry.Text)
		}
		if len(entry.Definitions) == 0 {
			t.Fatalf("No definitions found for entry: %v", entry)
		}
		if len(entry.Translations) == 0 {
			t.Fatalf("No translations found for entry: %v", entry)
		}
		if len(entry.Phonetics) == 0 {
			t.Fatalf("No phonetics found for entry: %v", entry)
		}
		if len(entry.DetailedTranslations) == 0 {
			t.Fatalf("No detailed translations found for entry: %v", entry)
		}
		if len(entry.ExamplesHTML) == 0 {
			t.Fatalf("No examples found for entry: %v", entry)
		}
	}
}

func TestWithoutOptions(t *testing.T) {
	source := google_translate_v2.New()

	entries, err := source.FetchEntries(context.Background(), "temple")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	if len(entries) == 0 {
		t.Fatalf("No entries found")
	}

	for _, entry := range entries {
		if entry.Text != "temple" {
			t.Fatalf("Expected entry text to be 'temple', but got '%s'", entry.Text)
		}
		if len(entry.Definitions) == 0 {
			t.Fatalf("No definitions found for entry: %v", entry)
		}
		if len(entry.Phonetics) == 0 {
			t.Fatalf("No phonetics found for entry: %v", entry)
		}
		if len(entry.ExamplesHTML) == 0 {
			t.Fatalf("No examples found for entry: %v", entry)
		}
	}
}

func TestSentence01(t *testing.T) {
	source := google_translate_v2.New(
		google_translate_v2.WithTranslationOptions(
			&google_translate_v2.TranslationOptions{
				From: google_translate_v2.English,
				To:   google_translate_v2.German,
				HL:   google_translate_v2.English,
				TLD:  "com",
			},
		),
	)

	entries, err := source.FetchEntries(context.Background(), "I will have gone by the time you get there")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	if len(entries) == 0 {
		t.Fatalf("No entries found")
	}

	for _, entry := range entries {
		if entry.Text != "I will have gone by the time you get there" {
			t.Fatalf("Expected entry text to be 'I will have gone by the time you get there', but got '%s'", entry.Text)
		}
		if len(entry.Translations) == 0 {
			t.Fatalf("No translations found for entry: %v", entry)
		}
	}
}

func TestIdiom(t *testing.T) {
	source := google_translate_v2.New(google_translate_v2.WithTranslationOptions(&google_translate_v2.TranslationOptions{
		From: google_translate_v2.English,
		To:   google_translate_v2.Turkish,
		HL:   google_translate_v2.English,
		TLD:  "com",
	}))

	entries, err := source.FetchEntries(context.Background(), "Come off it.")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	if len(entries) == 0 {
		t.Fatalf("No entries found")
	}

	for _, entry := range entries {
		if entry.Text != "Come off it." {
			t.Fatalf("Expected entry text to be 'Come off it.', but got '%s'", entry.Text)
		}
		if len(entry.Translations) == 0 {
			t.Fatalf("No translations found for entry: %v", entry)
		}
	}
}

func TestGenderedWord(t *testing.T) {
	source := google_translate_v2.New(
		google_translate_v2.WithTranslationOptions(
			&google_translate_v2.TranslationOptions{
				From: google_translate_v2.English,
				To:   google_translate_v2.Spanish,
				HL:   google_translate_v2.English,
				TLD:  "com",
			},
		),
	)

	entries, err := source.FetchEntries(context.Background(), "kid")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	if len(entries) == 0 {
		t.Fatalf("No entries found")
	}

	for _, entry := range entries {
		if entry.Text != "kid" {
			t.Fatalf("Expected entry text to be 'kid', but got '%s'", entry.Text)
		}
		if len(entry.Translations) == 0 {
			t.Fatalf("No translations found for entry: %v", entry)
		}
		for _, translation := range entry.Translations {
			if translation.Gender == "" {
				t.Fatalf("Expected translation gender not to be empty, but got '%s'", translation.Gender)
			}
		}
	}
}

func TestGenderedSentence(t *testing.T) {
	source := google_translate_v2.New(google_translate_v2.WithTranslationOptions(&google_translate_v2.TranslationOptions{
		From: google_translate_v2.English,
		To:   google_translate_v2.Spanish,
		HL:   google_translate_v2.English,
		TLD:  "com",
	}))

	entries, err := source.FetchEntries(context.Background(), "You left me perplexed")
	if err != nil {
		t.Fatalf("Error fetching word entries: %v", err)
	}

	if len(entries) == 0 {
		t.Fatalf("No entries found")
	}

	for _, entry := range entries {
		if entry.Text != "You left me perplexed" {
			t.Fatalf("Expected entry text to be 'You left me perplexed', but got '%s'", entry.Text)
		}
		if len(entry.Translations) == 0 {
			t.Fatalf("No translations found for entry: %v", entry)
		}
		for _, translation := range entry.Translations {
			if translation.Gender == "" {
				t.Fatalf("Expected translation gender not to be empty, but got '%s'", translation.Gender)
			}
		}
	}
}

func TestReadmeExample(t *testing.T) {
	translator := google_translate_v2.New(
		google_translate_v2.WithTranslationOptions(
			&google_translate_v2.TranslationOptions{
				From: google_translate_v2.English,
				To:   google_translate_v2.Spanish,
				HL:   google_translate_v2.English,
				TLD:  "com",
			},
		),
	)
	// Fetch translations
	entries, err := translator.FetchEntries(context.Background(), "Hello, world!")
	if err != nil {
		t.Fatalf("No entries found")
	}
	// Process entries
	for _, entry := range entries {
		if entry.Text == "" {
			t.Fatalf("Expected entry text not to be empty', but got '%s'", entry.Text)
		}
		if len(entry.Translations) == 0 {
			t.Fatalf("No translations found for entry: %v", entry)
		}
	}
}
