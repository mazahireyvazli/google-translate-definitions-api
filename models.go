package google_translate_v2

type Phonetic struct {
	Text string `json:"text"`
}

type Synonym struct {
	Text         string `json:"text"`
	UsageContext string `json:"usageContext"`
}

type Definition struct {
	Phonetics    []Phonetic `json:"phonetics"`
	Definition   string     `json:"definition"`
	Synonyms     []Synonym  `json:"synonyms"`
	Examples     []string   `json:"examples"`
	PartOfSpeech string     `json:"partOfSpeech"`
}

type Translation struct {
	Translation string `json:"translation"`
	Gender      string `json:"gender"`
}

type DetailedTranslation struct {
	Translation  string   `json:"translation"`
	Synonyms     []string `json:"synonyms"`
	PartOfSpeech string   `json:"partOfSpeech"`
	Frequency    float64  `json:"frequency"` // the lower the frequency, the more common the translation is
}

type Entry struct {
	Text                 string                `json:"text"`
	Phonetics            []Phonetic            `json:"phonetics"`
	Definitions          []Definition          `json:"definitions"`
	Translations         []Translation         `json:"translations"`
	DetailedTranslations []DetailedTranslation `json:"detailedTranslations"`
	ExamplesHTML         []string              `json:"examples"`
}
