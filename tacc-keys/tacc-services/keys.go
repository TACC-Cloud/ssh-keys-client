package services

import (
	"fmt"
)

// PublicKey reflects the structure of a successful response from the keys
// service.
type PublicKey struct {
	Message string          `json:"message"`
	Status  string          `json:"status"`
	Version string          `json:"version"`
	Result  ResultsResponse `json:"result"`
}

// ResultsResponse contains the fields of the "result" field from the key
// service's response.
type ResultsResponse struct {
	ID        int64         `json:"id"`
	PubKey    string        `json:"key_value"`
	Username  string        `json:"username"`
	Tenant    string        `json:"tenant"`
	CreatedAt string        `json:"created"`
	Tags      []ResultsTags `json:"tags"`
}

// ResultsTags holds any tags associated with the ResultsResponse struct.
type ResultsTags struct {
	Name string `json:"name"`
}

// String representation of a PublicKey.
func (p PublicKey) String() string {
	return fmt.Sprintf("%d %s\n%s", p.Result.ID, p.Result.Tags[0].Name, p.Result.PubKey)
}
