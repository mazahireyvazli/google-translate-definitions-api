# Google Translate V2

This package provides a Go implementation for interacting with Google Translate Public API.

## Features

- Customizable request options
- Configurable HTTP client
- Flexible parsing function

## Usage

```go
import "github.com/mazahireyvazli/google-translate-definitions-api"
// Create a new GoogleTranslateV2 instance with default options
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
// Handle error
}
// Process entries
for , entry := range entries {
// Use entry data
}
```

## Configuration

The `GoogleTranslateV2` struct can be configured using the following options:

- `WithTranslationOptions`: Set custom translation options
- `WithRequestOptions`: Set custom request options
- `WithHttpClient`: Use a custom HTTP client
- `WithParserFn`: Provide a custom parsing function

## Contacts

For questions, suggestions, or support, please contact:

- Email: mazahir.eyvazli@gmail.com
- GitHub: [mazahireyvazli](https://github.com/mazahireyvazli)

Feel free to open an issue on the GitHub repository if you encounter any problems or have feature requests.
