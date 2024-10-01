package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// full request logic
func makeRequest(requestType, url, token string, payload []byte) ([]byte, error) {
	client := &http.Client{}

	var request *http.Request
	if payload != nil {
		requestBody := bytes.NewReader(payload)
		request, _ = http.NewRequest(requestType, url, requestBody)
	} else {
		request, _ = http.NewRequest(requestType, url, nil)
	}

	request.Header.Set("Accept", "application/vnd.github+json")

	if token != "" {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	responce, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w, ", err)
	}

	body, _ := io.ReadAll(responce.Body)
	return body, nil

}

// send GET request
func MakeGetRequest(url, token string) ([]byte, error) {
	return makeRequest("GET", url, token, nil)
}

// send POST request
func MakePostRequest(url, token string, payload []byte) ([]byte, error) {
	return makeRequest("POST", url, token, payload)
}
