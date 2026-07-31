package client

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

type decodedModel struct {
	Name string `json:"name"`
}

type customDecodedModel struct {
	Value string
	Err   error
}

func (m *customDecodedModel) GetActualInstance() interface{} {
	return m
}

func (m *customDecodedModel) UnmarshalJSON(data []byte) error {
	if m.Err != nil {
		return m.Err
	}
	m.Value = string(data)
	return nil
}

type missingCustomUnmarshaller struct{}

func (m *missingCustomUnmarshaller) GetActualInstance() interface{} {
	return m
}

func TestDecodeJSONContentTypes(t *testing.T) {
	contentTypes := []string{
		"application/json",
		"application/hal+json",
		"application/problem+json",
	}

	for _, contentType := range contentTypes {
		t.Run(contentType, func(t *testing.T) {
			var got decodedModel
			err := newRequestTestClient(NewConfiguration()).Decode(
				&got,
				[]byte(`{"name":"example"}`),
				contentType,
			)
			if err != nil {
				t.Fatalf("Decode() error = %v", err)
			}
			if got.Name != "example" {
				t.Fatalf("decoded name = %q, want %q", got.Name, "example")
			}
		})
	}
}

func TestDecodeEmptyBodyDoesNotMutateDestination(t *testing.T) {
	got := decodedModel{Name: "unchanged"}
	err := newRequestTestClient(NewConfiguration()).Decode(&got, nil, "application/json")
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if got.Name != "unchanged" {
		t.Fatalf("decoded name = %q, want unchanged", got.Name)
	}
}

func TestDecodeStringBody(t *testing.T) {
	var got string
	err := newRequestTestClient(NewConfiguration()).Decode(
		&got,
		[]byte("plain response"),
		"application/octet-stream",
	)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if got != "plain response" {
		t.Fatalf("decoded string = %q, want %q", got, "plain response")
	}
}

func TestDecodeFailures(t *testing.T) {
	apiClient := newRequestTestClient(NewConfiguration())

	t.Run("malformed JSON", func(t *testing.T) {
		var got decodedModel
		if err := apiClient.Decode(&got, []byte(`{"name":`), "application/json"); err == nil {
			t.Fatal("Decode() error = nil, want malformed JSON error")
		}
	})

	t.Run("unsupported content type", func(t *testing.T) {
		var got decodedModel
		err := apiClient.Decode(&got, []byte("body"), "application/octet-stream")
		if err == nil || err.Error() != "undefined response type" {
			t.Fatalf("Decode() error = %v, want undefined response type", err)
		}
	})

	t.Run("custom unmarshal failure", func(t *testing.T) {
		wantErr := errors.New("custom failure")
		got := customDecodedModel{Err: wantErr}
		err := apiClient.Decode(&got, []byte(`{"name":"example"}`), "application/json")
		if !errors.Is(err, wantErr) {
			t.Fatalf("Decode() error = %v, want %v", err, wantErr)
		}
	})

	t.Run("missing custom unmarshaller", func(t *testing.T) {
		var got missingCustomUnmarshaller
		err := apiClient.Decode(&got, []byte(`{}`), "application/json")
		if err == nil || !strings.Contains(err.Error(), "no unmarshalObj.UnmarshalJSON defined") {
			t.Fatalf("Decode() error = %v, want missing custom unmarshaller error", err)
		}
	})
}

func TestDecodeCustomUnmarshaller(t *testing.T) {
	var got customDecodedModel
	err := newRequestTestClient(NewConfiguration()).Decode(
		&got,
		[]byte(`{"name":"example"}`),
		"application/json",
	)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if got.Value != `{"name":"example"}` {
		t.Fatalf("custom decoded value = %q, want raw JSON", got.Value)
	}
}

func TestDecodeFileBody(t *testing.T) {
	apiClient := newRequestTestClient(NewConfiguration())
	want := []byte{0x00, 0x01, 0x7f, 0x80, 0xfe, 0xff}

	var got *os.File
	if err := apiClient.Decode(&got, want, "application/octet-stream"); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if got == nil {
		t.Fatal("Decode() returned a nil file")
	}

	fileName := got.Name()
	t.Cleanup(func() {
		if err := got.Close(); err != nil {
			t.Errorf("close decoded file: %v", err)
		}
		if err := os.Remove(fileName); err != nil && !errors.Is(err, os.ErrNotExist) {
			t.Errorf("remove decoded file: %v", err)
		}
	})

	offset, err := got.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatalf("get decoded file offset: %v", err)
	}
	if offset != 0 {
		t.Fatalf("decoded file offset = %d, want 0", offset)
	}

	contents, err := io.ReadAll(got)
	if err != nil {
		t.Fatalf("read decoded file: %v", err)
	}
	if !bytes.Equal(contents, want) {
		t.Fatalf("decoded file contents = %v, want %v", contents, want)
	}
}
