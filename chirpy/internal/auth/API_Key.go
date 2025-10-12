package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")

	if auth == "" {
		return "", fmt.Errorf("Authorization token is empty")
	}
	// Expected format: "ApiKey <THE_KEY_HERE>"
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "ApiKey" {
		return "", fmt.Errorf("malformed authorization header, expected 'ApiKey <THE_KEY_HERE>'")
	}

	return parts[1], nil
}
