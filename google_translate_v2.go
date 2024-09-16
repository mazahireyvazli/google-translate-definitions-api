package google_translate_v2

import (
	"context"
	"net/http"
)

type RequestOptions struct {
	From    LangKey
	To      LangKey
	HL      LangKey
	TLD     string
	RPCIDs  string
	Headers map[string]string
}

type GoogleTranslateV2 struct {
	requestOptions *RequestOptions
	httpClient     *http.Client
	parserFn       func([]byte) ([]Entry, error)
}

type Option func(*GoogleTranslateV2)

func WithRequestOptions(ro *RequestOptions) Option {
	return func(gt *GoogleTranslateV2) {
		gt.requestOptions = ro
	}
}

func WithHttpClient(client *http.Client) Option {
	return func(gt *GoogleTranslateV2) {
		gt.httpClient = client
	}
}

func WithParserFn(fn func([]byte) ([]Entry, error)) Option {
	return func(gt *GoogleTranslateV2) {
		gt.parserFn = fn
	}
}

func New(options ...Option) *GoogleTranslateV2 {
	gt := &GoogleTranslateV2{
		requestOptions: &RequestOptions{
			From:    Auto,
			To:      English,
			HL:      English,
			TLD:     "com",
			RPCIDs:  "MkEWBc",
			Headers: make(map[string]string),
		},
		httpClient: &http.Client{},
		parserFn:   parseRawData,
	}

	for _, option := range options {
		if option != nil {
			option(gt)
		}
	}

	return gt
}

func (s *GoogleTranslateV2) FetchEntries(ctx context.Context, input string) ([]Entry, error) {
	data, err := sendRequest(ctx, input, s.requestOptions, s.httpClient)
	if err != nil {
		return nil, err
	}

	return s.parserFn(data)
}
