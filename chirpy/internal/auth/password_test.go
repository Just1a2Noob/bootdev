package auth

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "valid password",
			password: "validpassword123",
			wantErr:  false,
		},
		{
			name:     "empty password",
			password: "",
			wantErr:  false, // bcrypt can hash empty strings
		},
		{
			name:     "short password",
			password: "123",
			wantErr:  false,
		},
		{
			name:     "long password",
			password: strings.Repeat("a", 100),
			wantErr:  false,
		},
		{
			name:     "password with special characters",
			password: "p@ssw0rd!@#$%^&*()",
			wantErr:  false,
		},
		{
			name:     "unicode password",
			password: "пароль123",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)

			if tt.wantErr {
				if err == nil {
					t.Errorf("HashPassword() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("HashPassword() unexpected error: %v", err)
				return
			}

			// Verify hash is not empty
			if hash == "" {
				t.Errorf("HashPassword() returned empty hash")
			}

			// Verify hash starts with bcrypt prefix
			if !strings.HasPrefix(hash, "$2a$14$") {
				t.Errorf("HashPassword() hash doesn't have expected bcrypt prefix, got: %s", hash[:10])
			}

			// Verify we can validate the password against the hash
			err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(tt.password))
			if err != nil {
				t.Errorf("HashPassword() generated hash that doesn't validate against original password: %v", err)
			}
		})
	}
}

func TestHashPasswordConsistency(t *testing.T) {
	password := "testpassword"

	// Generate two hashes of the same password
	hash1, err1 := HashPassword(password)
	hash2, err2 := HashPassword(password)

	if err1 != nil {
		t.Errorf("HashPassword() first call failed: %v", err1)
	}
	if err2 != nil {
		t.Errorf("HashPassword() second call failed: %v", err2)
	}

	// Hashes should be different (bcrypt uses salt)
	if hash1 == hash2 {
		t.Errorf("HashPassword() generated identical hashes, expected different due to salt")
	}

	// Both hashes should validate against the original password
	if err := bcrypt.CompareHashAndPassword([]byte(hash1), []byte(password)); err != nil {
		t.Errorf("First hash doesn't validate: %v", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash2), []byte(password)); err != nil {
		t.Errorf("Second hash doesn't validate: %v", err)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	// Generate a valid hash for testing
	password := "testpassword123"
	validHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		t.Fatalf("Failed to generate test hash: %v", err)
	}

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "valid password and hash",
			password: password,
			hash:     string(validHash),
			wantErr:  false,
		},
		{
			name:     "wrong password",
			password: "wrongpassword",
			hash:     string(validHash),
			wantErr:  true,
		},
		{
			name:     "empty password with valid hash",
			password: "",
			hash:     string(validHash),
			wantErr:  true,
		},
		{
			name:     "valid password with invalid hash",
			password: password,
			hash:     "invalid_hash",
			wantErr:  true,
		},
		{
			name:     "empty hash",
			password: password,
			hash:     "",
			wantErr:  true,
		},
		{
			name:     "both empty",
			password: "",
			hash:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)

			if tt.wantErr {
				if err == nil {
					t.Errorf("CheckPasswordHash() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("CheckPasswordHash() unexpected error: %v", err)
				}
			}
		})
	}
}

func TestCheckPasswordHashWithGeneratedHash(t *testing.T) {
	testCases := []string{
		"password123",
		"",
		"very_long_password_with_special_characters_!@#$%^&*()",
		"短密码",
		"p@ssW0rd!",
	}

	for _, password := range testCases {
		t.Run("password_"+password, func(t *testing.T) {
			// Generate hash using our function
			hash, err := HashPassword(password)
			if err != nil {
				t.Fatalf("HashPassword() failed: %v", err)
			}

			// Verify correct password validates
			err = CheckPasswordHash(password, hash)
			if err != nil {
				t.Errorf("CheckPasswordHash() failed for correct password: %v", err)
			}

			// Verify wrong password fails
			wrongPassword := password + "wrong"
			err = CheckPasswordHash(wrongPassword, hash)
			if err == nil {
				t.Errorf("CheckPasswordHash() should have failed for wrong password")
			}
		})
	}
}

// Benchmark tests
func BenchmarkHashPassword(b *testing.B) {
	password := "benchmarkpassword123"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := HashPassword(password)
		if err != nil {
			b.Fatalf("HashPassword() failed: %v", err)
		}
	}
}

func BenchmarkCheckPasswordHash(b *testing.B) {
	password := "benchmarkpassword123"
	hash, err := HashPassword(password)
	if err != nil {
		b.Fatalf("Failed to generate test hash: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := CheckPasswordHash(password, hash)
		if err != nil {
			b.Fatalf("CheckPasswordHash() failed: %v", err)
		}
	}
}
