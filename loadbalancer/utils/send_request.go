package utils

import (
	"fmt"
	"io"
	"net/http"
)

// SendRequest posts the given request body to the provided URL and returns the
// response body. The caller is responsible for interpreting the returned data.
func SendRequest(url string, body io.Reader) ([]byte, error) {
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, nil
}
