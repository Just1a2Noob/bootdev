package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")

	if auth == "" {
		return "", fmt.Errorf("Authorization token is empty")
	}
	// Expected format: "Bearer <token>"
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("malformed authorization header, expected 'Bearer <token>'")
	}

	return parts[1], nil
}
