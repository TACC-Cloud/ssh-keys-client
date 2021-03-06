package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// TokenResponse represents a successful response from the Agave API upon
// requesting a refreshed token.
type TokenResponse struct {
	Scope        string `json:"scope"`
	TokeType     string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string
}

// RefreshToken makes an API call via HTTP to obtain a a new access and refresh
// token pair.
func RefreshToken(baseURL, refreshToken, apiKey, apiSecret string) (string, string, error) {
	// Token refresh endpoint.
	tokenEndpoint := baseURL + "/token"

	// Request data.
	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("scope", "PRODUCTION")
	v.Set("refresh_token", refreshToken)
	data := v.Encode()
	// Make request.
	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(data))
	if err != nil {
		return "", "", err
	}
	// Set request headers.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Basic HTTTP authentication.
	req.SetBasicAuth(apiKey, apiSecret)

	// Create HTTP client with timeout of 5s.
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	// Make HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// Parse in response.
	var refreshedToken TokenResponse
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&refreshedToken); err != nil {
			return "", "", err
		}

		// Add creation time.
		refreshedToken.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)

	} else { // API call returned an error.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", "", err
		}
		return "", "", errors.New(string(body))
	}

	return refreshedToken.AccessToken, refreshedToken.RefreshToken, nil
}

// GetTokens obtains an access and a refresh token pair.
func GetTokens(baseURL, apiKey, apiSecret, username, password string) (string, string, error) {
	// Token refresh endpoint.
	tokenEndpoint := baseURL + "/token"

	// Request data.
	v := url.Values{}
	v.Set("grant_type", "password")
	v.Set("scope", "PRODUCTION")
	v.Set("username", username)
	v.Set("password", password)
	data := v.Encode()
	// Make request.
	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(data))
	if err != nil {
		return "", "", err
	}
	// Set request headers.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Basic HTTTP authentication.
	req.SetBasicAuth(apiKey, apiSecret)

	// Create HTTP client with timeout of 5s.
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	// Make HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// Parse in response.
	var tokens TokenResponse
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
			return "", "", err
		}

		// Add creation time.
		tokens.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)

	} else { // API call returned an error.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", "", err
		}
		return "", "", errors.New(string(body))
	}
	return tokens.AccessToken, tokens.RefreshToken, nil
}
