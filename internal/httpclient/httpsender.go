package httpclient

import (
	"crypto/tls"
	"io"
	"net/http"
)

func NewRequest(req *http.Request) (string, error) {
	// Create the HTTP client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Ignore invalid certificates
		},
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != 200 {
		return "", err
	}

	// Get the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
