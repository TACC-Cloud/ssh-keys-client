package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type clientResponse struct {
	Result clientData `json:"result"`
}

type clientData struct {
	APIKey    string `json:"consumerKey"`
	APISecret string `json:"consumerSecret"`
}

// CreateClient creates a TACC oauth client with the provided name and
// description.
func CreateClient(baseURL, name, description, username, password string) (string, string, error) {
	// Oauth clients endpoint.
	clientEndpoint := baseURL + "/clients/v2"

	// Request data.
	v := url.Values{}
	v.Set("clientName", name)
	v.Set("tier", "Unlimited")
	v.Set("description", description)
	v.Set("callbackUrl", "")
	data := v.Encode()
	// Form request.
	req, err := http.NewRequest("POST", clientEndpoint, strings.NewReader(data))
	if err != nil {
		return "", "", err
	}

	// Set headers.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Basic http authentication.
	req.SetBasicAuth(username, password)

	// Create http client.
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	// Make request.
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		var client clientResponse
		if err := json.NewDecoder(resp.Body).Decode(&client); err != nil {
			return "", "", err
		}

		return client.Result.APIKey, client.Result.APISecret, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	return "", "", errors.New(string(body))
}
