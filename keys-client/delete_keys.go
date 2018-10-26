package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// DeletePubKey makes a request t the key server to delete the public key
// matching the given key id.
func (c *Configurations) DeletePubKey(keyId string) error {
	// Keys endpoint.
	keysEndpoint := c.BaseUrl + "/keys/v2/delete/" + keyId

	// Make request.
	req, err := http.NewRequest("DELETE", keysEndpoint, nil)
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
		fmt.Printf("key %s was successfully deleted\n", keyId)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(body))
	}

	return nil
}
