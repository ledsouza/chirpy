package auth

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	testCases := []struct {
		name        string
		userID      uuid.UUID
		secret      string
		expiresIn   time.Duration
		expectError bool
	}{
		{
			name:        "Valid token creation",
			userID:      uuid.New(),
			secret:      "test-secret",
			expiresIn:   time.Hour,
			expectError: false,
		},
		{
			name:        "Empty secret",
			userID:      uuid.New(),
			secret:      "",
			expiresIn:   time.Hour,
			expectError: false,
		},
		{
			name:        "Zero duration",
			userID:      uuid.New(),
			secret:      "test-secret",
			expiresIn:   0,
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := MakeJWT(tc.userID, tc.secret, tc.expiresIn)

			if tc.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Did not expect error but got: %v", err)
			}
			if err == nil {
				// Verify token structure
				if !strings.Contains(token, ".") {
					t.Error("Token doesn't contain dots separator")
				}

				// Verify token can be parsed
				parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
					return []byte(tc.secret), nil
				})
				if err != nil && tc.expiresIn != 0 {
					t.Errorf("Could not parse generated token: %v", err)
				}
				if err == nil && tc.expiresIn == 0 {
					t.Error("Expected error but got none")
				}
				if !parsed.Valid && tc.expiresIn != 0 {
					t.Error("Generated token is not valid")
				}
			}
		})
	}
}

func TestValidateJWT(t *testing.T) {
	testCases := []struct {
		name        string
		setup       func() (string, string, uuid.UUID)
		expectError bool
	}{
		{
			name: "Valid token",
			setup: func() (string, string, uuid.UUID) {
				userID := uuid.New()
				secret := "test-secret"
				token, _ := MakeJWT(userID, secret, time.Hour)
				return token, secret, userID
			},
			expectError: false,
		},
		{
			name: "Invalid secret",
			setup: func() (string, string, uuid.UUID) {
				userID := uuid.New()
				token, _ := MakeJWT(userID, "original-secret", time.Hour)
				return token, "wrong-secret", userID
			},
			expectError: true,
		},
		{
			name: "Expired token",
			setup: func() (string, string, uuid.UUID) {
				userID := uuid.New()
				secret := "test-secret"
				token, _ := MakeJWT(userID, secret, -time.Hour)
				return token, secret, userID
			},
			expectError: true,
		},
		{
			name: "Invalid token format",
			setup: func() (string, string, uuid.UUID) {
				return "invalid.token.format", "test-secret", uuid.Nil
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, secret, expectedID := tc.setup()
			userID, err := ValidateJWT(token, secret)

			if tc.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.expectError {
				if err != nil {
					t.Errorf("Did not expect error but got: %v", err)
				}
				if userID != expectedID {
					t.Errorf("Expected user ID %v but got %v", expectedID, userID)
				}
			}
		})
	}
}
func TestGetBearerToken(t *testing.T) {
	testCases := []struct {
		name        string
		header      string
		expectToken string
		expectError bool
	}{
		{
			name:        "Valid bearer token",
			header:      "Bearer abc123.def456.ghi789",
			expectToken: "abc123.def456.ghi789",
			expectError: false,
		},
		{
			name:        "Missing token",
			header:      "",
			expectToken: "",
			expectError: true,
		},
		{
			name:        "Invalid prefix",
			header:      "NotBearer abc123",
			expectToken: "",
			expectError: true,
		},
		{
			name:        "Bearer prefix only",
			header:      "Bearer ",
			expectToken: "",
			expectError: true,
		},
		{
			name:        "Missing space after Bearer",
			header:      "Bearerabc123",
			expectToken: "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			headers := make(http.Header)
			if tc.header != "" {
				headers.Add("Authorization", tc.header)
			}

			token, err := GetBearerToken(headers)

			if tc.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Did not expect error but got: %v", err)
			}
			if token != tc.expectToken {
				t.Errorf("Expected token %q but got %q", tc.expectToken, token)
			}
		})
	}
}
