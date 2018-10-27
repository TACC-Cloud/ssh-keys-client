package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SubscribeClient a TACC oauth client to TACC's public ssh keys service.
func SubscribeClient(baseURL, name, username, password string) error {
	// Oauth client ssubscription endpoint.
	clientEndpoint := baseURL + "/clients/v2/" + name + "/subscriptions"

	// Request data.
	v := url.Values{}
	v.Set("apiName", "PublicKeys")
	v.Set("apiVersion", "v2")
	v.Set("apiProvider", "admin")
	data := v.Encode()
	// Form request.
	req, err := http.NewRequest("POST", clientEndpoint, strings.NewReader(data))
	if err != nil {
		return err
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
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
