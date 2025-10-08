package sha256_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing/sha256"
)

func TestSha256_Sum(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected string
	}{
		{
			name:     "empty string",
			data:     []byte(""),
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "simple string",
			data:     []byte("hello"),
			expected: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			name:     "long string",
			data:     []byte("this is a longer test string for sha256"),
			expected: "b6359d59c087e0532bcab820157a3d5bc898165bc604d054d6fbb8b0b170aef7",
		},
	}

	hasher := sha256.NewSha256()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hasher.Sum(tt.data)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}
}
