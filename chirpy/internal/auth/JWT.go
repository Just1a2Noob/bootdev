package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn ...time.Duration) (string, error) {
	expiration := time.Hour
	if len(expiresIn) > 0 && expiresIn[0] > 0 {
		expiration = expiresIn[0]
	}
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  &jwt.NumericDate{time.Now()},
		ExpiresAt: &jwt.NumericDate{time.Now().Add(expiration)},
		Subject:   userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return fmt.Sprintf("Error signing token with secret : %s", err), err
	}

	return signed, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	// Apparently using an empty jwt.RegisteredClaims is the same as the function created before
	// Which means we just needed the structure

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// This is a security check to ensure the token was signed with the expected algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// The parser will use this key to verify the token's signature
		return []byte(tokenSecret), nil
	})

	// Handle parsing errors
	if err != nil {
		return uuid.Nil, fmt.Errorf("error parsing token: %w", err)
	}

	//Check if token is valid and claims are properly parsed
	if !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	// Validate the issuer (optional but recommended)
	if claims.Issuer != "chirpy" {
		return uuid.Nil, fmt.Errorf("invalid token issuer: %s", claims.Issuer)
	}

	user, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, err
	}

	return user, nil
}
