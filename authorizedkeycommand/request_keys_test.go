package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserPubKeys(t *testing.T) {
	// Mock host/<username>/text endpoint.
	m := http.NewServeMux()
	m.HandleFunc("/test/text", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "\n\nssh-rsa some-key-contents\n")
	}))
	ts := httptest.NewServer(m)
	defer ts.Close()

	// Set url.
	keysEndpoint = ts.URL

	var buf bytes.Buffer
	username := "test"
	if err := GetUserPubKeys(&buf, username); err != nil {
		t.Error(err)
	}
	t.Logf("\n'%s'\n", buf.String())
}
