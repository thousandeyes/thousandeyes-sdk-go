package utils

import "testing"

type strictModel struct {
	Name string `json:"name"`
}

func TestNewStrictDecoder(t *testing.T) {
	t.Run("known fields", func(t *testing.T) {
		var got strictModel
		if err := NewStrictDecoder([]byte(`{"name":"example"}`)).Decode(&got); err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
		if got.Name != "example" {
			t.Fatalf("Name = %q, want %q", got.Name, "example")
		}
	})

	t.Run("unknown field", func(t *testing.T) {
		var got strictModel
		if err := NewStrictDecoder([]byte(`{"name":"example","unknown":true}`)).Decode(&got); err == nil {
			t.Fatal("Decode() error = nil, want unknown-field error")
		}
	})

	t.Run("malformed JSON", func(t *testing.T) {
		var got strictModel
		if err := NewStrictDecoder([]byte(`{"name":`)).Decode(&got); err == nil {
			t.Fatal("Decode() error = nil, want malformed JSON error")
		}
	})
}
