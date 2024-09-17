package google_translate_v2

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func escapeSpecialSymbols(inputString string) string {
	escapedString := strings.ReplaceAll(inputString, `"`, `\\\\"`)
	normalizedString := strings.ReplaceAll(escapedString, "\n", "\\\\n")
	return normalizedString
}

func generateRequestBody(text string, translationOptions *TranslationOptions, requestOptions *RequestOptions) string {
	normalizedText := escapeSpecialSymbols(strings.TrimSpace(text))
	encodedData := url.QueryEscape(fmt.Sprintf(`[[["%s","[[\"%s\",\"%s\",\"%s\",1],[]]",null,"generic"]]]`, requestOptions.RPCIDs, normalizedText, translationOptions.From, translationOptions.To))
	return "f.req=" + encodedData + "&"
}

func generateRequestURL(translationOptions *TranslationOptions, requestOptions *RequestOptions) (string, error) {
	params := url.Values{}
	params.Add("rpcids", requestOptions.RPCIDs)
	params.Add("source-path", "/")
	params.Add("hl", string(translationOptions.HL))
	params.Add("soc-app", "1")
	params.Add("soc-platform", "1")
	params.Add("soc-device", "1")

	return fmt.Sprintf("https://translate.google.%s/_/TranslateWebserverUi/data/batchexecute?%s", translationOptions.TLD, params.Encode()), nil
}

func sendRequest(ctx context.Context, text string, translationOptions *TranslationOptions, requestOptions *RequestOptions, httpClient *http.Client) ([]byte, error) {
	body := generateRequestBody(text, translationOptions, requestOptions)
	url, err := generateRequestURL(translationOptions, requestOptions)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range requestOptions.Headers {
		req.Header.Set(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
