package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// DeletePubKey makes a request t the key server to delete the public key
// matching the given key id.
func DeletePubKey(baseURL, accessToken, keyID string) error {
	// Keys endpoint.
	keysEndpoint := baseURL + "/keys/v2/delete/" + keyID

	// Make request.
	req, err := http.NewRequest("DELETE", keysEndpoint, nil)
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
		fmt.Printf("key %s was successfully deleted\n", keyID)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}
