package hs256_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto/hs256"
)

func TestJWT(t *testing.T) {
	tests := []struct {
		name      string
		key       []byte
		data      crypto.DataClaims
		expire    time.Time
		tokenEdit func(string) string
		wantErr   bool
	}{
		{
			name:    "valid token",
			key:     []byte("secret"),
			data:    crypto.DataClaims{},
			expire:  time.Now().Add(time.Hour),
			wantErr: false,
		},
		{
			name:    "invalid signature",
			key:     []byte("wrong-secret"),
			data:    crypto.DataClaims{},
			expire:  time.Now().Add(time.Hour),
			wantErr: true,
		},
		{
			name:    "expired token",
			key:     []byte("secret"),
			data:    crypto.DataClaims{},
			expire:  time.Now().Add(-time.Minute),
			wantErr: true,
		},
		{
			name:      "malformed token",
			key:       []byte("secret"),
			tokenEdit: func(_ string) string { return "invalid.token.string" },
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		j := hs256.NewJWT([]byte("secret"))
		var token string
		var err error

		if tt.tokenEdit != nil {
			token = tt.tokenEdit("")
		} else {
			token, err = j.Generate(tt.data, tt.expire)
			if err != nil {
				t.Fatalf("%s: generate error: %v", tt.name, err)
			}
			if tt.name == "invalid signature" {
				j = hs256.NewJWT(tt.key)
			}
		}

		got, err := j.Validate(token)
		if (err != nil) != tt.wantErr {
			t.Fatalf("%s: unexpected error state, got %v", tt.name, err)
		}
		if !tt.wantErr && !reflect.DeepEqual(got, tt.data) {
			t.Fatalf("%s: claims mismatch, got %#v want %#v", tt.name, got, tt.data)
		}
	}
}
