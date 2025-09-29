package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Test the happy path - creating and validating a valid token
func TestValidateJWT_ValidToken(t *testing.T) {
	// Arrange
	userID := uuid.New()
	secret := "test-secret-key"
	expiresIn := time.Hour

	// Act - Create token
	tokenString, err := makeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("makeJWT failed: %v", err)
	}

	// Act - Validate token
	validatedUserID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}

	// Assert
	if validatedUserID != userID {
		t.Errorf("Expected userID %v, got %v", userID, validatedUserID)
	}
}

// Test validation with wrong secret
func TestValidateJWT_WrongSecret(t *testing.T) {
	// Arrange
	userID := uuid.New()
	correctSecret := "correct-secret"
	wrongSecret := "wrong-secret"
	expiresIn := time.Hour

	// Act - Create token with correct secret
	tokenString, err := makeJWT(userID, correctSecret, expiresIn)
	if err != nil {
		t.Fatalf("makeJWT failed: %v", err)
	}

	// Act - Try to validate with wrong secret
	_, err = ValidateJWT(tokenString, wrongSecret)

	// Assert - Should fail
	if err == nil {
		t.Error("Expected validation to fail with wrong secret, but it succeeded")
	}
}

// Test validation of expired token
func TestValidateJWT_ExpiredToken(t *testing.T) {
	// Arrange
	userID := uuid.New()
	secret := "test-secret"
	// Create token that expires immediately
	expiresIn := -time.Hour // Already expired

	// Act - Create expired token
	tokenString, err := makeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("makeJWT failed: %v", err)
	}

	// Act - Try to validate expired token
	_, err = ValidateJWT(tokenString, secret)

	// Assert - Should fail
	if err == nil {
		t.Error("Expected validation to fail for expired token, but it succeeded")
	}
}

// Test validation with malformed token
func TestValidateJWT_MalformedToken(t *testing.T) {
	// Arrange
	secret := "test-secret"
	malformedTokens := []string{
		"not.a.jwt",
		"invalid-token",
		"",
		"header.payload", // Missing signature
		"too.many.parts.here.invalid",
	}

	for _, tokenString := range malformedTokens {
		t.Run("malformed_token_"+tokenString, func(t *testing.T) {
			// Act
			_, err := ValidateJWT(tokenString, secret)

			// Assert - Should fail
			if err == nil {
				t.Errorf("Expected validation to fail for malformed token '%s', but it succeeded", tokenString)
			}
		})
	}
}

// Test validation with tampered token
func TestValidateJWT_TamperedToken(t *testing.T) {
	// Arrange
	userID := uuid.New()
	secret := "test-secret"
	expiresIn := time.Hour

	// Act - Create valid token
	tokenString, err := makeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("makeJWT failed: %v", err)
	}

	// Tamper with the token by changing one character
	tamperedToken := tokenString[:len(tokenString)-1] + "X"

	// Act - Try to validate tampered token
	_, err = ValidateJWT(tamperedToken, secret)

	// Assert - Should fail
	if err == nil {
		t.Error("Expected validation to fail for tampered token, but it succeeded")
	}
}

// Test with different signing algorithm (should fail)
func TestValidateJWT_WrongSigningAlgorithm(t *testing.T) {
	// Arrange
	userID := uuid.New()
	secret := "test-secret"

	// Create a token with a different algorithm (this is a bit tricky to test directly)
	// We'll create a token manually with HS512 instead of HS256
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour)},
		Subject:   userID.String(),
	}

	// Create token with HS512 instead of HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Failed to create HS512 token: %v", err)
	}

	// Act - Try to validate (our ValidateJWT expects HS256)
	_, err = ValidateJWT(tokenString, secret)

	// Assert - Should fail due to algorithm mismatch
	if err == nil {
		t.Error("Expected validation to fail for different signing algorithm, but it succeeded")
	}
}

// Test with empty secret
func TestMakeJWT_EmptySecret(t *testing.T) {
	// Arrange
	userID := uuid.New()
	emptySecret := ""
	expiresIn := time.Hour

	// Act
	tokenString, err := makeJWT(userID, emptySecret, expiresIn)

	// Assert - Should work (empty secret is technically valid for HMAC)
	if err != nil {
		t.Errorf("makeJWT failed with empty secret: %v", err)
	}

	// Validate the token with the same empty secret
	validatedUserID, err := ValidateJWT(tokenString, emptySecret)
	if err != nil {
		t.Errorf("ValidateJWT failed with empty secret: %v", err)
	}

	if validatedUserID != userID {
		t.Errorf("Expected userID %v, got %v", userID, validatedUserID)
	}
}

// Test with zero UUID
func TestValidateJWT_ZeroUUID(t *testing.T) {
	// Arrange
	userID := uuid.Nil // Zero UUID
	secret := "test-secret"
	expiresIn := time.Hour

	// Act
	tokenString, err := makeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("makeJWT failed: %v", err)
	}

	validatedUserID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}

	// Assert
	if validatedUserID != uuid.Nil {
		t.Errorf("Expected zero UUID %v, got %v", uuid.Nil, validatedUserID)
	}
}

// Test token claims validation (issuer check)
func TestValidateJWT_WrongIssuer(t *testing.T) {
	// Arrange
	userID := uuid.New()
	secret := "test-secret"

	// Create a token with different issuer
	claims := jwt.RegisteredClaims{
		Issuer:    "wrong-issuer", // Different from expected "chirpy"
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour)},
		Subject:   userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	// Act
	_, err = ValidateJWT(tokenString, secret)

	// Assert - Should fail due to wrong issuer
	if err == nil {
		t.Error("Expected validation to fail for wrong issuer, but it succeeded")
	}
}

// Test token without subject
func TestValidateJWT_MissingSubject(t *testing.T) {
	// Arrange
	secret := "test-secret"

	// Create a token without subject
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour)},
		// Subject is missing
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	// Act
	_, err = ValidateJWT(tokenString, secret)

	// Assert - Should fail due to missing subject
	if err == nil {
		t.Error("Expected validation to fail for missing subject, but it succeeded")
	}
}

// Test token with invalid UUID in subject
func TestValidateJWT_InvalidUUIDInSubject(t *testing.T) {
	// Arrange
	secret := "test-secret"

	// Create a token with invalid UUID in subject
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour)},
		Subject:   "not-a-valid-uuid",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	// Act
	_, err = ValidateJWT(tokenString, secret)

	// Assert - Should fail due to invalid UUID
	if err == nil {
		t.Error("Expected validation to fail for invalid UUID, but it succeeded")
	}
}

// Benchmark test for performance
func BenchmarkMakeJWT(b *testing.B) {
	userID := uuid.New()
	secret := "test-secret"
	expiresIn := time.Hour

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := makeJWT(userID, secret, expiresIn)
		if err != nil {
			b.Fatalf("makeJWT failed: %v", err)
		}
	}
}

func BenchmarkValidateJWT(b *testing.B) {
	userID := uuid.New()
	secret := "test-secret"
	expiresIn := time.Hour

	tokenString, err := makeJWT(userID, secret, expiresIn)
	if err != nil {
		b.Fatalf("makeJWT failed: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ValidateJWT(tokenString, secret)
		if err != nil {
			b.Fatalf("ValidateJWT failed: %v", err)
		}
	}
}
