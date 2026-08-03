package thousandeyes_test

import (
	"encoding/json"
	"testing"

	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
)

func TestGeneratedEnumJSONCompatibility(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    administrative.CloudEnterpriseAgentType
		wantErr bool
	}{
		{
			name:  "known value",
			input: `"cloud"`,
			want:  administrative.CLOUDENTERPRISEAGENTTYPE_CLOUD,
		},
		{
			name:  "explicit unknown value",
			input: `"unknown"`,
			want:  administrative.CLOUDENTERPRISEAGENTTYPE_UNKNOWN,
		},
		{
			name:  "future value falls back",
			input: `"future-agent-type"`,
			want:  administrative.CLOUDENTERPRISEAGENTTYPE_UNKNOWN,
		},
		{
			name:    "malformed JSON",
			input:   `"unterminated`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got administrative.CloudEnterpriseAgentType
			err := json.Unmarshal([]byte(test.input), &got)
			if test.wantErr {
				if err == nil {
					t.Fatal("json.Unmarshal() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			if got != test.want {
				t.Fatalf("decoded enum = %q, want %q", got, test.want)
			}
			if !got.IsValid() {
				t.Fatalf("decoded enum %q is not valid", got)
			}
		})
	}
}
