package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// refreshTokenResponse represents a successful response from the Agave API upon
// requesting a refreshed token.
type refreshTokenResponse struct {
	Scope        string `json:"scope"`
	TokeType     string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// responseError holds the values of a failed request.
type responseError struct {
	Type        string `json:"error"`
	Description string `json:"error_description"`
	Status      string
}

// Error formats uses the fields of the responseError struct to format an error
// message.
func (e *responseError) Error() string {
	return fmt.Sprintf("%s: %s - %s", e.Status, e.Type, e.Description)
}

// RefreshAPIToken makes an API call via HTTP to obtain a a new access and
// refresh token pair.
func (c *Configurations) RefreshAPIToken() error {
	// Token refresh endpoint.
	tokenEndpoint := c.BaseUrl + "/token"

	// Request data.
	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("scope", "PRODUCTION")
	v.Set("refresh_token", c.RefreshToken)
	data := v.Encode()
	// Make request.
	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(data))
	if err != nil {
		fmt.Printf("Error building request: %s\n", err)
		return err
	}
	// Set request headers.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Basic HTTTP authentication.
	req.SetBasicAuth(c.ApiKey, c.ApiSecret)

	// Create HTTP client with timeout of 10s.
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Make HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return err
	}
	defer resp.Body.Close()

	// Check request was successful before proceeding to update configurations
	// file.
	if resp.StatusCode == http.StatusOK {
		// Take the response's body and store it into a refreshTokenBody
		// struct.
		var refreshedToken refreshTokenResponse
		if err := json.NewDecoder(resp.Body).Decode(&refreshedToken); err != nil {
			fmt.Printf("Error decoding response body: %s\n", err)
			return err
		}

		// Update refresh and access token.
		c.RefreshToken = refreshedToken.RefreshToken
		c.AccessToken = refreshedToken.AccessToken
		c.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)
		c.ExpiresIn = strconv.FormatInt(refreshedToken.ExpiresIn, 10)

		// Update configurations file.
		if err := c.SaveConfig(); err != nil {
			fmt.Printf("Error writing to config file: %s\n", err)
			return err
		}
	} else { // API call returned an error.
		var failedResp responseError
		if err := json.NewDecoder(resp.Body).Decode(&failedResp); err != nil {
			fmt.Printf("Error decoding failed response's body: %s\n", err)
			return err
		}
		failedResp.Status = resp.Status
		return &failedResp
	}

	return nil
}
