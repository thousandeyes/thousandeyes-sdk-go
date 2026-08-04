package error_test

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	internalerror "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/error"
)

type problemDetails struct {
	Title  *string `json:"title,omitempty"`
	Detail *string `json:"detail,omitempty"`
}

type messageDetails struct {
	Message string `json:"message"`
}

func (p *problemDetails) GetTitleOk() (*string, bool) {
	if p == nil || p.Title == nil {
		return nil, false
	}
	return p.Title, true
}

func (p *problemDetails) GetDetailOk() (*string, bool) {
	if p == nil || p.Detail == nil {
		return nil, false
	}
	return p.Detail, true
}

func TestDecodeErrorFormatsModeledMessages(t *testing.T) {
	apiClient := &client.APIClient{}
	tests := []struct {
		name         string
		body         string
		model        interface{}
		wantMessage  string
		wantModel    interface{}
		contentType  string
		responseCode int
	}{
		{
			name:         "RFC7807 title and detail",
			body:         `{"title":"Access denied","detail":"The token lacks permission"}`,
			model:        &problemDetails{},
			wantMessage:  "403 Forbidden Access denied (The token lacks permission)",
			wantModel:    problemDetails{Title: stringPointer("Access denied"), Detail: stringPointer("The token lacks permission")},
			contentType:  "application/problem+json",
			responseCode: http.StatusForbidden,
		},
		{
			name:         "RFC7807 title only",
			body:         `{"title":"Access denied"}`,
			model:        &problemDetails{},
			wantMessage:  "403 Forbidden Access denied",
			wantModel:    problemDetails{Title: stringPointer("Access denied")},
			contentType:  "application/problem+json",
			responseCode: http.StatusForbidden,
		},
		{
			name:         "RFC7807 detail only",
			body:         `{"detail":"The token lacks permission"}`,
			model:        &problemDetails{},
			wantMessage:  "403 Forbidden The token lacks permission",
			wantModel:    problemDetails{Detail: stringPointer("The token lacks permission")},
			contentType:  "application/problem+json",
			responseCode: http.StatusForbidden,
		},
		{
			name:         "RFC7807 neither title nor detail uses trimmed raw body",
			body:         "\n  {\"type\":\"https://example.com/problems/forbidden\",\"status\":403} \t",
			model:        &problemDetails{},
			wantMessage:  `403 Forbidden {"type":"https://example.com/problems/forbidden","status":403}`,
			wantModel:    problemDetails{},
			contentType:  "application/problem+json",
			responseCode: http.StatusForbidden,
		},
		{
			name:         "non-RFC model falls back to raw body",
			body:         `{"message":"Access denied"}`,
			model:        &messageDetails{},
			wantMessage:  `403 Forbidden {"message":"Access denied"}`,
			wantModel:    messageDetails{Message: "Access denied"},
			contentType:  "application/json",
			responseCode: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := []byte(tt.body)
			response := errorResponse(tt.responseCode, tt.contentType)

			apiErr := internalerror.DecodeError(apiClient.Decode, tt.model, body, response)

			if got := apiErr.ErrorMessage; got != tt.wantMessage {
				t.Errorf("GenericAPIError.ErrorMessage = %q, want %q", got, tt.wantMessage)
			}
			if !bytes.Equal(apiErr.Body, body) {
				t.Errorf("GenericAPIError.Body = %q, want %q", apiErr.Body, body)
			}
			if apiErr.Model != tt.model {
				t.Errorf("GenericAPIError.Model = %p, want original model %p", apiErr.Model, tt.model)
			}
			if got := reflect.ValueOf(apiErr.Model).Elem().Interface(); !reflect.DeepEqual(got, tt.wantModel) {
				t.Errorf("GenericAPIError.Model value = %#v, want %#v", got, tt.wantModel)
			}
		})
	}
}

func TestDecodeErrorPreservesDecodeFailuresAndEmptyBodies(t *testing.T) {
	apiClient := &client.APIClient{}
	tests := []struct {
		name        string
		body        string
		contentType string
		wantMessage string
		wantModel   bool
	}{
		{
			name:        "malformed JSON",
			body:        `{"title":`,
			contentType: "application/problem+json",
			wantMessage: "unexpected end of JSON input",
		},
		{
			name:        "non-JSON response",
			body:        "gateway exploded",
			contentType: "text/plain",
			wantMessage: "undefined response type",
		},
		{
			name:        "empty modeled response",
			contentType: "application/problem+json",
			wantMessage: "403 Forbidden",
			wantModel:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := []byte(tt.body)
			model := &problemDetails{}

			apiErr := internalerror.DecodeError(
				apiClient.Decode,
				model,
				body,
				errorResponse(http.StatusForbidden, tt.contentType),
			)

			if got := apiErr.ErrorMessage; got != tt.wantMessage {
				t.Errorf("GenericAPIError.ErrorMessage = %q, want %q", got, tt.wantMessage)
			}
			if !bytes.Equal(apiErr.Body, body) {
				t.Errorf("GenericAPIError.Body = %q, want %q", apiErr.Body, body)
			}
			if tt.wantModel && apiErr.Model != model {
				t.Errorf("GenericAPIError.Model = %p, want original model %p", apiErr.Model, model)
			}
			if !tt.wantModel && apiErr.Model != nil {
				t.Errorf("GenericAPIError.Model = %#v, want nil", apiErr.Model)
			}
		})
	}
}

func TestDecodeErrorSkipsDecodeForUnmodeledResponse(t *testing.T) {
	body := []byte(`{"message":"short and stout"}`)
	decodeCalled := false
	decode := func(interface{}, []byte, string) error {
		decodeCalled = true
		return nil
	}

	apiErr := internalerror.DecodeError(
		decode,
		nil,
		body,
		errorResponse(http.StatusTeapot, "application/json"),
	)

	if decodeCalled {
		t.Error("decode was called for an unmodeled response")
	}
	if got, want := apiErr.ErrorMessage, "418 I'm a teapot"; got != want {
		t.Errorf("GenericAPIError.ErrorMessage = %q, want %q", got, want)
	}
	if !bytes.Equal(apiErr.Body, body) {
		t.Errorf("GenericAPIError.Body = %q, want %q", apiErr.Body, body)
	}
	if apiErr.Model != nil {
		t.Errorf("GenericAPIError.Model = %#v, want nil", apiErr.Model)
	}
}

func errorResponse(statusCode int, contentType string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Status:     fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode)),
		Header: http.Header{
			"Content-Type": []string{contentType},
		},
	}
}

func stringPointer(value string) *string {
	return &value
}
