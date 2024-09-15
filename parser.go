package google_translate_v2

import (
	"encoding/json"
	"fmt"
	"strings"
)

var typeOfSpeeches = []string{
	"Noun", "Verb", "Adjective", "Adverb", "Preposition",
	"Abbreviation", "Conjunction", "Pronoun", "Interjection",
	"Phrase", "Prefix", "Suffix", "Article", "Combining form",
	"Numeral", "Auxiliary verb", "Exclamation", "Plural", "Particle",
}

func getTypeOfSpeech(typeNumber int) string {
	if typeNumber < 1 || typeNumber > len(typeOfSpeeches) {
		return "generic"
	}
	return typeOfSpeeches[typeNumber-1]
}

func parseRawData(rawData []byte) ([]WordEntry, error) {
	rawDataStr := string(rawData)
	rawDataStr = strings.TrimPrefix(rawDataStr, ")]}'")

	var rawBodyTyped [][]any
	if err := json.Unmarshal([]byte(rawDataStr), &rawBodyTyped); err != nil {
		return nil, fmt.Errorf("error unmarshalling rawDataStr: %w", err)
	}

	entriesPortion := rawBodyTyped[0][2].(string)
	var rawObj []any
	if err := json.Unmarshal([]byte(entriesPortion), &rawObj); err != nil {
		return nil, fmt.Errorf("error unmarshalling entriesPortion: %w", err)
	}

	var phoneticText string
	var actualWord string

	if len(rawObj) > 0 {
		if slice1, ok := rawObj[0].([]any); ok && len(slice1) > 1 {
			if pt, ok := slice1[0].(string); ok {
				phoneticText = pt
			}

			// Check for actualWord in multiple locations
			if slice2, ok := slice1[1].([]any); ok && len(slice2) > 0 {
				if slice3, ok := slice2[0].([]any); ok {
					if len(slice3) > 4 {
						if str, ok := slice3[4].(string); ok {
							actualWord = str
						}
					}
					if actualWord == "" && len(slice3) > 1 {
						if str, ok := slice3[1].(string); ok {
							actualWord = str
						}
					}
				}
			}

			// Check rawObj[0][6][0] if actualWord is still empty
			if actualWord == "" && len(slice1) > 6 {
				if slice6, ok := slice1[6].([]any); ok && len(slice6) > 0 {
					if str, ok := slice6[0].(string); ok {
						actualWord = str
					}
				}
			}
		}
	}

	var definitions []Definition

	if len(rawObj) > 3 {
		if slice3, ok := rawObj[3].([]any); ok && len(slice3) > 1 {
			if slice31, ok := slice3[1].([]any); ok && len(slice31) > 0 {
				if slice310, ok := slice31[0].([]any); ok {
					for _, item := range slice310 {
						itemSlice, ok := item.([]any)
						if !ok || len(itemSlice) < 2 {
							continue
						}

						var partOfSpeech string

						if len(itemSlice) >= 4 {
							if typeNumber, ok := itemSlice[3].(float64); ok {
								partOfSpeech = getTypeOfSpeech(int(typeNumber))
							}
						}

						if definitionsSlice, ok := itemSlice[1].([]any); ok {
							for _, definitionElementSlice := range definitionsSlice {

								definitionElementSlice, ok := definitionElementSlice.([]any)
								if !ok || len(definitionElementSlice) == 0 {
									continue
								}

								var definition string
								var examples []string

								if definitionStr, ok := definitionElementSlice[0].(string); ok {
									definition = definitionStr
								}

								if len(definitionElementSlice) > 1 {
									if exampleStr, ok := definitionElementSlice[1].(string); ok {
										if len(exampleStr) > 0 {
											examples = append(examples, exampleStr)
										}
									}
								}

								synonyms := make(map[string][]string)
								if len(definitionElementSlice) > 5 {
									if synonymsSlice, ok := definitionElementSlice[5].([]any); ok {
										for _, synonym := range synonymsSlice {
											if synonymObj, ok := synonym.([]any); ok && len(synonymObj) > 0 {
												var synonymType string = "normal"
												if len(synonymObj) > 1 {
													if synonymTypeObj, ok := synonymObj[1].([]any); ok {
														synonymType = synonymTypeObj[0].([]any)[0].(string)
													}
												}

												if synonymStringsSlice, ok := synonymObj[0].([]any); ok {
													for _, synonymStringSliceItem := range synonymStringsSlice {
														if synonymStringSliceItemObj, ok := synonymStringSliceItem.([]any); ok && len(synonymStringSliceItemObj) > 0 {
															if synonymString, ok := synonymStringSliceItemObj[0].(string); ok {
																synonyms[synonymType] = append(synonyms[synonymType], synonymString)
															}
														}

													}
												}
											}
										}
									}
								}

								definitions = append(definitions, Definition{
									PartOfSpeech: partOfSpeech,
									Definition:   definition,
									Synonyms:     synonyms,
									Antonyms:     []string{},
									Examples:     examples,
								})

							}
						}
					}
				}
			}
		}
	}

	if len(definitions) == 0 {
		return nil, fmt.Errorf("no definitions found for word: %s", actualWord)
	}

	wordEntry := WordEntry{
		Word:        actualWord,
		Definitions: definitions,
		Phonetics: []Phonetic{
			{
				Text: phoneticText,
			},
		},
	}

	return []WordEntry{wordEntry}, nil
}
