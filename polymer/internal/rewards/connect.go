package rewards

import (
	"fmt"
	"net/http"
	"time"
)

const (
	APIUrl         = "https://rewardscc-api.azure-api.net/v1"
	APIKey         = "a33f53b560b747e9a29bb7b6be234a45"
	RequestTimeout = 10 * time.Second
)

var Endpoints = map[string]string{
	"card_list":   "creditcard-cardlist",
	"card_detail": "creditcard-detail-bycard",
}

// FetchEndpoint makes an authenticated GET request to the specified endpoint.
func FetchEndpoint(endpointName string, params []string) (*http.Response,
	error) {

	endpoint, exists := Endpoints[endpointName]
	if !exists {
		return nil, fmt.Errorf("endpoint '%s' does not exist", endpointName)
	}

	url := fmt.Sprintf("%s/%s", APIUrl, endpoint)
	for _, param := range params {
		url += "/" + param
	}
	url += "?skey=" + APIKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{
		Timeout: RequestTimeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	} else if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp, nil
}
