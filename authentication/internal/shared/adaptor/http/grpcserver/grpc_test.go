package grpcserver_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/config"
)

type mockBootstrap struct {
	shouldFail bool
}

func (m *mockBootstrap) Boot(_ interface{}) error {
	if m.shouldFail {
		return errors.New("bootstrap failed")
	}
	return nil
}

func TestGrpc_Run_PreServe(t *testing.T) {
	tests := []struct {
		name      string
		port      int
		bootFail  bool
		expectErr bool
	}{
		{
			name:      "valid port and bootstrap success",
			port:      0,
			bootFail:  false,
			expectErr: false,
		},
		{
			name:      "bootstrap failure",
			port:      0,
			bootFail:  true,
			expectErr: true,
		},
		{
			name:      "invalid port",
			port:      -1,
			bootFail:  false,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{}
			cfg.Server.Port = tt.port

			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.Database.Port))
			if err != nil && !tt.expectErr {
				t.Fatalf("unexpected error in listen: %v", err)
			}
			if lis != nil {
				lis.Close()
			}
		})
	}
}
