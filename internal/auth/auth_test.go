package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantKey   string
		expectErr bool
	}{
		{
			name:      "missing Authorization header",
			headers:   http.Header{},
			wantKey:   "",
			expectErr: true,
		},
		{
			name: "empty Authorization header",
			headers: http.Header{
				"Authorization": {""},
			},
			wantKey:   "",
			expectErr: true,
		},
		{
			name: "invalid format - missing ApiKey",
			headers: http.Header{
				"Authorization": {"Bearer xyz123"},
			},
			wantKey:   "",
			expectErr: true,
		},
		{
			name: "invalid format - missing key value",
			headers: http.Header{
				"Authorization": {"ApiKey"},
			},
			wantKey:   "",
			expectErr: true,
		},
		{
			name: "valid ApiKey header",
			headers: http.Header{
				"Authorization": {"ApiKey my-api-key"},
			},
			wantKey:   "my-api-key",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)

			if (err != nil) != tt.expectErr {
				t.Fatalf("expected error: %v, got error: %v", tt.expectErr, err)
			}
			if gotKey != tt.wantKey {
				t.Errorf("expected key: %s, got key: %s", tt.wantKey, gotKey)
			}
		})
	}
}
