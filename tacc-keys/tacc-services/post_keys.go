package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type payload struct {
	KeyValue string        `json:"key_value"`
	Tags     []payloadTags `json:"tags"`
}

type payloadTags struct {
	Purpose string `json:"name"`
}

// PostUserPubKey posts a user's public key to the keys server.
func PostUserPubKey(baseURL, accessToken, user, pubkey string) error {
	// Keys endpoint.
	keysEndpoint := baseURL + "/keys/v2/" + user

	// Request payload.
	data := payload{
		KeyValue: pubkey,
		Tags:     []payloadTags{{Purpose: "go-keys-client"}},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Make request.
	req, err := http.NewRequest("POST", keysEndpoint, bytes.NewReader(payloadBytes))
	if err != nil {
		return err
	}

	// Set request headers.
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

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

		return errors.New(string(body))
	}

	return nil
}
