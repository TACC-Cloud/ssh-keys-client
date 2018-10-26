package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Request payload
type Payload struct {
	KeyValue string        `json:"key_value"`
	Tags     []PayloadTags `json:"tags"`
}

type PayloadTags struct {
	Purpose string `json:"name"`
}

// PostUserPubKey posts a user's public key to the keys server.
func (c *Configurations) PostUserPubKey(user string, pubkey string) error {
	// Keys endpoint.
	keysEndpoint := c.BaseUrl + "/keys/v2/" + user

	// Request payload.
	data := Payload{
		KeyValue: pubkey,
		Tags:     []PayloadTags{{Purpose: "go-keys-client"}},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Make request.
	req, err := http.NewRequest("POST", keysEndpoint, bytes.NewReader(payloadBytes))
	if err != nil {
		fmt.Printf("Error building request: %s\n", err)
		return err
	}

	// Set request headers.
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

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
		var pubkey PublicKey
		if err := json.NewDecoder(resp.Body).Decode(&pubkey); err != nil {
			return err
		}

		fmt.Println(pubkey)
	} else { // API call returned an error.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(body))
	}

	return nil
}
