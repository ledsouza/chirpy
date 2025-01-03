package auth

import "testing"

func TestCheckPasswordHash(t *testing.T) {
	passwordHash, err := HashPassword("correct-password")
	if err != nil {
		t.Errorf("error hashing password: %s", err.Error())
		return
	}

	tests := []struct {
		name        string
		password    string
		hash        string
		shouldMatch bool
	}{
		{
			name:        "valid password and hash should match",
			password:    "correct-password",
			hash:        passwordHash, // hash for "correct-password"
			shouldMatch: true,
		},
		{
			name:        "invalid password should not match hash",
			password:    "wrong-password",
			hash:        "$2a$10$OhI55Jm/8qXmzjHERAYXSuZNDZ8KdOH5rYaHoYLGzkGhr0L8rNRqa",
			shouldMatch: false,
		},
		{
			name:        "invalid hash format should fail",
			password:    "password",
			hash:        "invalid-hash",
			shouldMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if tt.shouldMatch && err != nil {
				t.Errorf("expected password to match hash, got error: %v", err)
			}
			if !tt.shouldMatch && err == nil {
				t.Error("expected password not to match hash, but it did")
			}
		})
	}
}
