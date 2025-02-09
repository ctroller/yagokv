package client

import (
	"fmt"
	"io"
	"net/http"
)

// API_ENDPOINT is the base API path used for key-value operations.
var API_ENDPOINT = "/api/v1/kvs/"

// RemoteClient represents the HTTP client with a base URL.
type RemoteClient struct {
	baseUrl string
}

// NewClient creates and returns a new Client instance with the given base URL.
func NewRemoteClient(baseUrl string) Client {
	return &RemoteClient{baseUrl: baseUrl}
}

// Get sends an HTTP GET request to retrieve the value associated with the key.
func (c *RemoteClient) Get(key string) (string, error) {
	return c.makeRequest(http.MethodGet, API_ENDPOINT, key)
}

// GetAndTransform retrieves the value for the specified key and applies the provided transformer.
//
// Parameters:
//   - key: Identifier for the key-value pair.
//   - transformer: Transformer to apply to the value.
func (c *RemoteClient) GetAndTransform(key string, transformer Transformer[string, any]) (any, error) {
	value, err := c.Get(key)
	if err != nil {
		return nil, err
	}

	return transformer.Transform(value), nil
}

// Delete sends an HTTP DELETE request to remove the resource identified by key.
func (c *RemoteClient) Delete(key string) error {
	_, err := c.makeRequest(http.MethodDelete, API_ENDPOINT, key)
	return err
}

// Set stores the key-value pair using an HTTP request.
//
// Parameters:
//   - key: Identifier for the key-value pair.
//   - value: Value to be stored.
func (c *RemoteClient) Set(key string, value string) error {
	_, err := c.makeRequest(http.MethodPost, API_ENDPOINT, key)
	return err
}

// getRequestUrl constructs the full URL for an API endpoint and key.
//
// Parameters:
//   - endpoint: The API endpoint path.
//   - key: Identifier or parameter to be appended to the endpoint.
//
// Returns:
//
//	The complete URL string.
func (c *RemoteClient) getRequestUrl(endpoint string, key string) string {
	return fmt.Sprintf("%s%s%s", c.baseUrl, endpoint, key)
}

// makeRequest creates and sends an HTTP request using the specified method,
// and returns the response body as a string or an error.
//
// Parameters:
//   - method: HTTP method to use (e.g., GET, DELETE).
//   - endpoint: The API endpoint path.
//   - key: Identifier to be appended to the endpoint and used in the URL.
func (c *RemoteClient) makeRequest(method string, endpoint string, key string) (string, error) {
	requestURL := c.getRequestUrl(endpoint, key)
	request, err := http.NewRequest(method, requestURL, nil)
	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
