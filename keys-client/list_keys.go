package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetUserPubKeys makes a request for plain-text public keys for a given user.
func (c *Configurations) GetUserPubKeys(user string) error {
	// Keys endpoint.
	keysEndpoint := c.BaseUrl + "/keys/v2/" + user

	// Make request.
	req, err := http.NewRequest("GET", keysEndpoint, nil)
	if err != nil {
		fmt.Printf("Error building request: %s\n", err)
		return err
	}

	// Set request headers.
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

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

	// Check if request was successful.
	if resp.StatusCode == http.StatusOK {
		var publicKeys []PublicKey
		// Pass public keys to stdout.
		if err := json.NewDecoder(resp.Body).Decode(&publicKeys); err != nil {
			return err
		}

		for _, pubkey := range publicKeys {
			fmt.Println(pubkey)
		}

	} else { // API call returned an error.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(body))
	}

	return nil
}
