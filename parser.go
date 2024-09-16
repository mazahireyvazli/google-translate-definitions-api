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

func parseRawData(rawData []byte) ([]Entry, error) {
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
	var actualEntry string
	var translations []Translation
	var detailedTranslations []DetailedTranslation
	var examples []string

	if len(rawObj) > 0 {
		if slice1, ok := rawObj[0].([]any); ok && len(slice1) > 1 {
			if pt, ok := slice1[0].(string); ok {
				phoneticText = pt
			}

			if slice2, ok := slice1[1].([]any); ok && len(slice2) > 0 {
				if slice3, ok := slice2[0].([]any); ok {
					if len(slice3) > 4 {
						if str, ok := slice3[4].(string); ok {
							actualEntry = str
						}
					}
					if actualEntry == "" && len(slice3) > 1 {
						if str, ok := slice3[1].(string); ok {
							actualEntry = str
						}
					}
				}
			}

			if actualEntry == "" && len(slice1) > 6 {
				if slice6, ok := slice1[6].([]any); ok && len(slice6) > 0 {
					if str, ok := slice6[0].(string); ok {
						actualEntry = str
					}
				}
			}
		}
	}

	if len(rawObj) > 1 {
		if translationsSlice, ok := rawObj[1].([]any); ok && len(translationsSlice) > 0 {
			if translationData, ok := translationsSlice[0].([]any); ok {

				if slice00, ok := translationData[0].([]any); ok && len(slice00) > 5 {
					if slice005, ok := slice00[5].([]any); ok && len(slice005) > 0 {
						if slice0050, ok := slice005[0].([]any); ok && len(slice0050) > 0 {
							if translatedText, ok := slice0050[0].(string); ok {
								translations = append(translations, Translation{Translation: translatedText})
							}
						}
					}
				}

				for _, translation := range translationData {
					if translationSlice, ok := translation.([]any); ok && len(translationSlice) > 2 {
						translatedText, textOk := translationSlice[0].(string)
						gender, genderOk := translationSlice[2].(string)

						if textOk {
							newTranslation := Translation{
								Translation: translatedText,
							}
							if genderOk {
								newTranslation.Gender = gender
							}
							translations = append(translations, newTranslation)
						}
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
								var synonyms []Synonym

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
																synonyms = append(synonyms, Synonym{
																	Text:         synonymString,
																	UsageContext: synonymType,
																})
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
									Examples:     examples,
								})

							}
						}
					}
				}
			}

			if len(slice3) > 2 {
				if slice32, ok := slice3[2].([]any); ok && len(slice32) > 0 {
					if slice320, ok := slice32[0].([]any); ok && len(slice320) > 0 {
						for _, item := range slice320 {
							if itemSlice, ok := item.([]any); ok && len(itemSlice) > 0 {
								if itemSlice1, ok := itemSlice[1].(string); ok {
									examples = append(examples, itemSlice1)
								}
							}
						}
					}
				}
			}

			if len(slice3) > 5 {
				if slice35, ok := slice3[5].([]any); ok && len(slice35) > 0 {
					if slice350, ok := slice35[0].([]any); ok && len(slice350) > 0 {
						for _, item := range slice350 {
							if itemSlice, ok := item.([]any); ok && len(itemSlice) > 0 {

								partOfSpeech := getTypeOfSpeech(int(itemSlice[4].(float64)))

								if translation, ok := itemSlice[1].([]any); ok && len(translation) > 0 {
									for _, translationItem := range translation {
										if translationItemSlice, ok := translationItem.([]any); ok && len(translationItemSlice) > 0 {
											detailedTranslation := DetailedTranslation{
												PartOfSpeech: partOfSpeech,
											}

											if translationItemStr, ok := translationItemSlice[0].(string); ok {
												detailedTranslation.Translation = translationItemStr
											}

											if synonyms, ok := translationItemSlice[2].([]any); ok && len(synonyms) > 0 {
												for _, synonym := range synonyms {
													if synonymStr, ok := synonym.(string); ok {
														detailedTranslation.Synonyms = append(detailedTranslation.Synonyms, synonymStr)
													}
												}
											}

											if frequency, ok := translationItemSlice[3].(float64); ok {
												detailedTranslation.Frequency = frequency
											}

											detailedTranslations = append(detailedTranslations, detailedTranslation)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return []Entry{{
		Text:        actualEntry,
		Definitions: definitions,
		Phonetics: []Phonetic{
			{
				Text: phoneticText,
			},
		},
		Translations:         translations,
		DetailedTranslations: detailedTranslations,
		ExamplesHTML:         examples,
	}}, nil
}
