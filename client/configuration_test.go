package client

import "testing"

func TestNewConfiguration(t *testing.T) {
	cfg := NewConfiguration()

	if cfg.ServerURL != "https://api.thousandeyes.com/v7" {
		t.Fatalf("ServerURL = %q, want %q", cfg.ServerURL, "https://api.thousandeyes.com/v7")
	}
	if cfg.Debug {
		t.Fatal("Debug = true, want false")
	}
}

func TestConfigurationFluentOptions(t *testing.T) {
	cfg := NewConfiguration()

	if got := cfg.WithAuthToken("test-token"); got != cfg {
		t.Fatal("WithAuthToken returned a different configuration")
	}
	if cfg.AuthToken != "test-token" {
		t.Fatalf("AuthToken = %q, want %q", cfg.AuthToken, "test-token")
	}

	if got := cfg.WithServerUrl("https://example.test/v7"); got != cfg {
		t.Fatal("WithServerUrl returned a different configuration")
	}
	if cfg.ServerURL != "https://example.test/v7" {
		t.Fatalf("ServerURL = %q, want %q", cfg.ServerURL, "https://example.test/v7")
	}
}

func TestBuildUserAgent(t *testing.T) {
	tests := []struct {
		name      string
		userAgent string
		want      string
	}{
		{
			name: "SDK only",
			want: "ThousandEyesSDK-Go/" + SDKVersion(),
		},
		{
			name:      "custom suffix",
			userAgent: "example-client/1.0",
			want:      "ThousandEyesSDK-Go/" + SDKVersion() + " example-client/1.0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := NewConfiguration()
			cfg.UserAgent = test.userAgent

			if got := cfg.BuildUserAgent(); got != test.want {
				t.Fatalf("BuildUserAgent() = %q, want %q", got, test.want)
			}
		})
	}
}
