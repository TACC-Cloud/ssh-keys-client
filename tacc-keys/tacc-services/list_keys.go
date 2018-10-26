package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetUserPubKeys makes a request for plain-text public keys for a given user.
func GetUserPubKeys(baseURL, accessToken, user string) error {
	// Keys endpoint.
	keysEndpoint := baseURL + "/keys/v2/" + user

	// Make request.
	req, err := http.NewRequest("GET", keysEndpoint, nil)
	if err != nil {
		return err
	}

	// Set request headers.
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Create HTTP client with timeout of 10s.
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	// Make HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if request was successful.
	if resp.StatusCode == http.StatusOK {
		var publicKeys []PublicKey
		if err := json.NewDecoder(resp.Body).Decode(&publicKeys); err != nil {
			return err
		}

		// Write keys to stdout.
		for _, pubkey := range publicKeys {
			fmt.Println(pubkey)
		}

	} else { // API call returned an error.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}
