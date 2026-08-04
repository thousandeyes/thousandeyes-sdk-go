package error_test

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/thousandeyes/thousandeyes-sdk-go/v3/bgpmonitors"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	internalerror "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/error"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func TestGenericAPIErrorErrorReturnsConfiguredMessage(t *testing.T) {
	apiErr := internalerror.GenericAPIError{
		ErrorMessage: "403 Forbidden access denied",
	}

	if got, want := apiErr.Error(), apiErr.ErrorMessage; got != want {
		t.Fatalf("GenericAPIError.Error() = %q, want %q", got, want)
	}
}

func TestGeneratedAPIErrorDecoding(t *testing.T) {
	tests := []struct {
		name             string
		statusCode       int
		contentType      string
		body             string
		wantErrorMessage string
		assertError      func(*testing.T, *internalerror.GenericAPIError)
	}{
		{
			name:        "modeled RFC7807 response preserves body and decodes model",
			statusCode:  http.StatusForbidden,
			contentType: "application/problem+json",
			body:        `{"type":"https://example.com/problems/forbidden","title":"Access denied","status":403,"detail":"The token lacks permission","instance":"/monitors"}`,
			assertError: func(t *testing.T, apiErr *internalerror.GenericAPIError) {
				t.Helper()

				model, ok := apiErr.Model.(*bgpmonitors.Error)
				if !ok {
					t.Fatalf("GenericAPIError.Model has type %T, want *bgpmonitors.Error", apiErr.Model)
				}

				want := bgpmonitors.Error{}
				want.SetType("https://example.com/problems/forbidden")
				want.SetTitle("Access denied")
				want.SetStatus(http.StatusForbidden)
				want.SetDetail("The token lacks permission")
				want.SetInstance("/monitors")
				if !reflect.DeepEqual(*model, want) {
					t.Errorf("GenericAPIError.Model = %#v, want %#v", *model, want)
				}
			},
		},
		{
			name:             "modeled RFC7807 response without title or detail uses trimmed raw body",
			statusCode:       http.StatusForbidden,
			contentType:      "application/problem+json",
			body:             "\n  {\"type\":\"https://example.com/problems/forbidden\",\"status\":403} \t",
			wantErrorMessage: `403 Forbidden {"type":"https://example.com/problems/forbidden","status":403}`,
			assertError: func(t *testing.T, apiErr *internalerror.GenericAPIError) {
				t.Helper()

				model, ok := apiErr.Model.(*bgpmonitors.Error)
				if !ok {
					t.Fatalf("GenericAPIError.Model has type %T, want *bgpmonitors.Error", apiErr.Model)
				}

				want := bgpmonitors.Error{}
				want.SetType("https://example.com/problems/forbidden")
				want.SetStatus(http.StatusForbidden)
				if !reflect.DeepEqual(*model, want) {
					t.Errorf("GenericAPIError.Model = %#v, want %#v", *model, want)
				}
			},
		},
		{
			name:             "modeled non-RFC JSON preserves compatibility message and model",
			statusCode:       http.StatusUnauthorized,
			contentType:      "application/json",
			body:             `{"error":"unauthorized","error_description":"token expired"}`,
			wantErrorMessage: `401 Unauthorized {"error":"unauthorized","error_description":"token expired"}`,
			assertError: func(t *testing.T, apiErr *internalerror.GenericAPIError) {
				t.Helper()

				model, ok := apiErr.Model.(*bgpmonitors.UnauthorizedError)
				if !ok {
					t.Fatalf("GenericAPIError.Model has type %T, want *bgpmonitors.UnauthorizedError", apiErr.Model)
				}

				want := bgpmonitors.UnauthorizedError{}
				want.SetError("unauthorized")
				want.SetErrorDescription("token expired")
				if !reflect.DeepEqual(*model, want) {
					t.Errorf("GenericAPIError.Model = %#v, want %#v", *model, want)
				}
			},
		},
		{
			name:             "malformed modeled JSON preserves decoder error and omits model",
			statusCode:       http.StatusForbidden,
			contentType:      "application/problem+json",
			body:             `{"title":`,
			wantErrorMessage: "unexpected end of JSON input",
			assertError: func(t *testing.T, apiErr *internalerror.GenericAPIError) {
				t.Helper()
				if apiErr.Model != nil {
					t.Errorf("GenericAPIError.Model = %#v, want nil", apiErr.Model)
				}
			},
		},
		{
			name:             "non-JSON modeled response preserves decoder error and omits model",
			statusCode:       http.StatusForbidden,
			contentType:      "text/plain",
			body:             "gateway exploded",
			wantErrorMessage: "undefined response type",
			assertError: func(t *testing.T, apiErr *internalerror.GenericAPIError) {
				t.Helper()
				if apiErr.Model != nil {
					t.Errorf("GenericAPIError.Model = %#v, want nil", apiErr.Model)
				}
			},
		},
		{
			name:             "empty modeled response uses status and zero-value model",
			statusCode:       http.StatusForbidden,
			contentType:      "application/problem+json",
			body:             "",
			wantErrorMessage: "403 Forbidden",
			assertError: func(t *testing.T, apiErr *internalerror.GenericAPIError) {
				t.Helper()

				model, ok := apiErr.Model.(*bgpmonitors.Error)
				if !ok {
					t.Fatalf("GenericAPIError.Model has type %T, want *bgpmonitors.Error", apiErr.Model)
				}
				if !reflect.DeepEqual(*model, bgpmonitors.Error{}) {
					t.Errorf("GenericAPIError.Model = %#v, want zero-value bgpmonitors.Error", *model)
				}
			},
		},
		{
			name:             "unmodeled response uses status and omits model",
			statusCode:       http.StatusTeapot,
			contentType:      "application/json",
			body:             `{"message":"short and stout"}`,
			wantErrorMessage: "418 I'm a teapot",
			assertError: func(t *testing.T, apiErr *internalerror.GenericAPIError) {
				t.Helper()
				if apiErr.Model != nil {
					t.Errorf("GenericAPIError.Model = %#v, want nil", apiErr.Model)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiErr := executeBGPMonitorsErrorResponse(
				t,
				tt.statusCode,
				tt.contentType,
				tt.body,
			)

			if got, want := string(apiErr.Body), tt.body; got != want {
				t.Errorf("GenericAPIError.Body = %q, want %q", got, want)
			}
			if tt.wantErrorMessage != "" && apiErr.ErrorMessage != tt.wantErrorMessage {
				t.Errorf("GenericAPIError.ErrorMessage = %q, want %q", apiErr.ErrorMessage, tt.wantErrorMessage)
			}
			tt.assertError(t, apiErr)
		})
	}
}

func executeBGPMonitorsErrorResponse(
	t *testing.T,
	statusCode int,
	contentType string,
	body string,
) *internalerror.GenericAPIError {
	t.Helper()

	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if got, want := request.Method, http.MethodGet; got != want {
			t.Errorf("request method = %q, want %q", got, want)
		}
		if got, want := request.URL.Path, "/monitors"; got != want {
			t.Errorf("request path = %q, want %q", got, want)
		}

		return &http.Response{
			StatusCode: statusCode,
			Status:     fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode)),
			Header: http.Header{
				"Content-Type": []string{contentType},
			},
			Body:    io.NopCloser(strings.NewReader(body)),
			Request: request,
		}, nil
	})

	cfg := client.NewConfiguration().WithServerUrl("https://example.invalid")
	cfg.HTTPClient = &http.Client{Transport: transport}
	apiClient := client.NewAPIClient(cfg)
	api := (*bgpmonitors.BGPMonitorsAPIService)(&apiClient.Common)

	_, _, err := api.GetBgpMonitors().Execute()
	if err == nil {
		t.Fatal("GetBgpMonitors().Execute() returned nil error, want *GenericAPIError")
	}

	apiErr, ok := err.(*internalerror.GenericAPIError)
	if !ok {
		t.Fatalf("GetBgpMonitors().Execute() error has type %T, want *GenericAPIError", err)
	}
	return apiErr
}
