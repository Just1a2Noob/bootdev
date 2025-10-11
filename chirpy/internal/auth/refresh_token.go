package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshTokens() (string, error) {
	// Refresh tokens are used to give permission to refresh the JWT token
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	hex_enc := hex.EncodeToString(key)

	return hex_enc, nil
}
