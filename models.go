package google_translate_v2

type PhoneticType string

const (
	PhoneticTypeUK          PhoneticType = "UK"
	PhoneticTypeUS          PhoneticType = "US"
	PhoneticTypeAU          PhoneticType = "AU"
	PhoneticTypeUNSPECIFIED PhoneticType = "UNSPECIFIED"
)

type Phonetic struct {
	Text      string       `json:"text"`
	Audio     string       `json:"audio"`
	SourceURL string       `json:"sourceUrl"`
	Type      PhoneticType `json:"type"`
}

type Definition struct {
	Phonetics    []Phonetic          `json:"phonetics"`
	Definition   string              `json:"definition"`
	Synonyms     map[string][]string `json:"synonyms"`
	Antonyms     []string            `json:"antonyms"`
	Examples     []string            `json:"examples"`
	PartOfSpeech string              `json:"partOfSpeech"`
}

type WordEntry struct {
	Word        string       `json:"word"`
	Phonetics   []Phonetic   `json:"phonetics"`
	Definitions []Definition `json:"definitions"`
}
