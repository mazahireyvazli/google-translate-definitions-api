package google_translate_v2

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func validateTLD(tld string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z]{2,63}$`, tld)
	return match
}

func escapeSpecialSymbols(inputString string) string {
	escapedString := strings.ReplaceAll(inputString, `"`, `\\\\"`)
	normalizedString := strings.ReplaceAll(escapedString, "\n", "\\\\n")
	return normalizedString
}

func generateRequestBody(text string, options *RequestOptions) string {
	normalizedText := escapeSpecialSymbols(strings.TrimSpace(text))
	encodedData := url.QueryEscape(fmt.Sprintf(`[[["%s","[[\"%s\",\"%s\",\"%s\",1],[]]",null,"generic"]]]`, options.RPCIDs, normalizedText, options.From, options.To))
	return "f.req=" + encodedData + "&"
}

func generateRequestURL(options *RequestOptions) (string, error) {
	params := url.Values{}
	params.Add("rpcids", options.RPCIDs)
	params.Add("source-path", "/")
	params.Add("hl", string(options.HL))
	params.Add("soc-app", "1")
	params.Add("soc-platform", "1")
	params.Add("soc-device", "1")

	return fmt.Sprintf("https://translate.google.%s/_/TranslateWebserverUi/data/batchexecute?%s", options.TLD, params.Encode()), nil
}

func sendRequest(ctx context.Context, text string, options *RequestOptions, httpClient *http.Client) ([]byte, error) {
	if !validateTLD(options.TLD) {
		return nil, fmt.Errorf("invalid TLD: Must be 2-63 letters only")
	}

	body := generateRequestBody(text, options)
	url, err := generateRequestURL(options)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range options.Headers {
		req.Header.Set(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
