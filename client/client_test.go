package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

type contextKey struct{}

func newRequestTestClient(cfg *Configuration) *APIClient {
	return &APIClient{cfg: cfg}
}

func TestDumpRedactedRequest(t *testing.T) {
	const requestBody = `{"password":"alpha beta","apiKey":"gamma&delta",` +
		`"token":"epsilon\"zeta","clientSecret":"eta,theta","name":"body-visible"}`
	request, err := http.NewRequest(
		http.MethodPost,
		"https://example.test/resource/path-visible?apiKey=query-secret&filter=query-visible",
		strings.NewReader(requestBody),
	)
	if err != nil {
		t.Fatalf("creating request: %v", err)
	}
	request.Header.Set("Authorization", "Bearer header-token")
	request.Header.Set("Cookie", "session=cookie-secret")
	request.Header.Set("X-Api-Key", "custom-header-secret")
	request.Header.Set("X-Trace", "header-visible")

	dump, err := dumpRedactedRequest(request)
	if err != nil {
		t.Fatalf("dumpRedactedRequest() error = %v", err)
	}
	logged := string(dump)

	for _, secret := range []string{
		"header-token",
		"cookie-secret",
		"custom-header-secret",
		"query-secret",
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
	} {
		if strings.Contains(logged, secret) {
			t.Errorf("debug request log contains sensitive value %q", secret)
		}
	}
	for _, detail := range []string{"path-visible", "query-visible", "header-visible", "body-visible"} {
		if !strings.Contains(logged, detail) {
			t.Errorf("debug request log does not contain request detail %q", detail)
		}
	}
	if !strings.Contains(logged, "Authorization: [REDACTED]") {
		t.Errorf("debug request log does not contain a redacted Authorization header\nlog:\n%s", logged)
	}

	if got := request.Header.Get("Authorization"); got != "Bearer header-token" {
		t.Errorf("original Authorization header = %q, want unchanged", got)
	}
	if got := request.Header.Get("Cookie"); got != "session=cookie-secret" {
		t.Errorf("original Cookie header = %q, want unchanged", got)
	}
	if got := request.Header.Get("X-Api-Key"); got != "custom-header-secret" {
		t.Errorf("original X-Api-Key header = %q, want unchanged", got)
	}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		t.Fatalf("reading original request body: %v", err)
	}
	if got := string(body); got != requestBody {
		t.Errorf("original request body = %q, want %q", got, requestBody)
	}
}

func TestRedactSensitiveResponseData(t *testing.T) {
	dump := []byte("HTTP/1.1 200 OK\r\n" +
		"Set-Cookie: session=response-cookie\r\n" +
		"X-Trace: response-header\r\n" +
		"Content-Type: application/json\r\n\r\n" +
		`{"refreshToken":"response token","result":"response-visible"}`)

	logged := string(redactSensitiveData(dump))
	for _, secret := range []string{"response-cookie", "response token"} {
		if strings.Contains(logged, secret) {
			t.Errorf("debug response log contains sensitive value %q", secret)
		}
	}
	for _, detail := range []string{"response-header", "response-visible"} {
		if !strings.Contains(logged, detail) {
			t.Errorf("debug response log does not contain response detail %q", detail)
		}
	}
}

func TestPrepareRequestHeadersQueryAndContext(t *testing.T) {
	baseCtx, cancel := context.WithCancel(context.Background())
	ctx := context.WithValue(baseCtx, contextKey{}, "context-value")
	cancel()
	cfg := NewConfiguration().WithAuthToken("synthetic-token")
	cfg.UserAgent = "test-suite/1.0"
	cfg.Context = ctx
	apiClient := newRequestTestClient(cfg)

	query := url.Values{
		"filter[name]": {"first", "second"},
		"new":          {"value"},
	}
	request, err := apiClient.PrepareRequest(
		"https://example.test/resource?existing=preserved",
		http.MethodGet,
		nil,
		map[string]string{"X-Test": "header"},
		query,
		nil,
	)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	if got := request.Header.Get("Authorization"); got != "Bearer synthetic-token" {
		t.Fatalf("Authorization = %q, want %q", got, "Bearer synthetic-token")
	}
	wantUserAgent := "ThousandEyesSDK-Go/" + SDKVersion() + " test-suite/1.0"
	if got := request.Header.Get("User-Agent"); got != wantUserAgent {
		t.Fatalf("User-Agent = %q, want %q", got, wantUserAgent)
	}
	if got := request.Header.Get("X-Test"); got != "header" {
		t.Fatalf("X-Test = %q, want %q", got, "header")
	}
	if got := request.URL.Query().Get("existing"); got != "preserved" {
		t.Fatalf("existing query = %q, want %q", got, "preserved")
	}
	if got := request.URL.Query()["filter[name]"]; len(got) != 2 || got[0] != "first" || got[1] != "second" {
		t.Fatalf("filter[name] query = %#v, want [first second]", got)
	}
	if !strings.Contains(request.URL.RawQuery, "filter[name]=") {
		t.Fatalf("RawQuery = %q, want unescaped bracketed key", request.URL.RawQuery)
	}
	if got := request.Context().Value(contextKey{}); got != "context-value" {
		t.Fatalf("request context value = %v, want %q", got, "context-value")
	}
	if err := request.Context().Err(); err != context.Canceled {
		t.Fatalf("request context error = %v, want %v", err, context.Canceled)
	}
}

func TestPrepareRequestOmitsAuthorizationWithoutToken(t *testing.T) {
	request, err := newRequestTestClient(NewConfiguration()).PrepareRequest(
		"https://example.test/resource",
		http.MethodGet,
		nil,
		map[string]string{},
		nil,
		nil,
	)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}
	if got := request.Header.Get("Authorization"); got != "" {
		t.Fatalf("Authorization = %q, want empty", got)
	}
}

func TestPrepareRequestJSONBody(t *testing.T) {
	headers := map[string]string{}
	request, err := newRequestTestClient(NewConfiguration()).PrepareRequest(
		"https://example.test/resource",
		http.MethodPost,
		map[string]string{"name": "example"},
		headers,
		nil,
		nil,
	)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	if got := request.Header.Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("Content-Type = %q, want JSON", got)
	}
	var body map[string]string
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		t.Fatalf("decoding request body: %v", err)
	}
	if got := body["name"]; got != "example" {
		t.Fatalf("body name = %q, want %q", got, "example")
	}
}

func TestPrepareRequestFormBody(t *testing.T) {
	form := url.Values{"name": {"example"}, "tag": {"one", "two"}}
	request, err := newRequestTestClient(NewConfiguration()).PrepareRequest(
		"https://example.test/resource",
		http.MethodPost,
		nil,
		map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
		nil,
		form,
	)
	if err != nil {
		t.Fatalf("PrepareRequest() error = %v", err)
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		t.Fatalf("reading request body: %v", err)
	}
	if got, want := string(body), form.Encode(); got != want {
		t.Fatalf("form body = %q, want %q", got, want)
	}
	if got, want := request.Header.Get("Content-Length"), "28"; got != want {
		t.Fatalf("Content-Length = %q, want %q", got, want)
	}
}

func TestPrepareRequestFailures(t *testing.T) {
	apiClient := newRequestTestClient(NewConfiguration())

	tests := []struct {
		name        string
		path        string
		method      string
		postBody    interface{}
		headers     map[string]string
		form        url.Values
		wantMessage string
	}{
		{
			name:        "body and form conflict",
			path:        "https://example.test/resource",
			method:      http.MethodPost,
			postBody:    "body",
			headers:     map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
			form:        url.Values{"name": {"example"}},
			wantMessage: "cannot specify postBody and x-www-form-urlencoded form at the same time",
		},
		{
			name:        "unsupported body",
			path:        "https://example.test/resource",
			method:      http.MethodPost,
			postBody:    make(chan int),
			headers:     map[string]string{},
			wantMessage: "invalid body type",
		},
		{
			name:        "JSON encoding failure",
			path:        "https://example.test/resource",
			method:      http.MethodPost,
			postBody:    make(chan int),
			headers:     map[string]string{"Content-Type": "application/json"},
			wantMessage: "unsupported type",
		},
		{
			name:        "invalid URL",
			path:        "https://example.test/%",
			method:      http.MethodGet,
			headers:     map[string]string{},
			wantMessage: "invalid URL escape",
		},
		{
			name:        "invalid method",
			path:        "https://example.test/resource",
			method:      "NOT VALID",
			headers:     map[string]string{},
			wantMessage: "invalid method",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := apiClient.PrepareRequest(
				test.path,
				test.method,
				test.postBody,
				test.headers,
				nil,
				test.form,
			)
			if err == nil {
				t.Fatal("PrepareRequest() error = nil, want error")
			}
			if !strings.Contains(err.Error(), test.wantMessage) {
				t.Fatalf("PrepareRequest() error = %q, want substring %q", err, test.wantMessage)
			}
		})
	}
}

func TestSetBodySupportedTypes(t *testing.T) {
	text := "pointer"
	tests := []struct {
		name        string
		body        interface{}
		contentType string
		want        string
	}{
		{name: "reader", body: bytes.NewBufferString("reader"), contentType: "text/plain", want: "reader"},
		{name: "bytes", body: []byte("bytes"), contentType: "application/octet-stream", want: "bytes"},
		{name: "string", body: "string", contentType: "text/plain", want: "string"},
		{name: "string pointer", body: &text, contentType: "text/plain", want: "pointer"},
		{name: "JSON", body: map[string]string{"name": "value"}, contentType: "application/json", want: "{\"name\":\"value\"}\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := setBody(test.body, test.contentType)
			if err != nil {
				t.Fatalf("setBody() error = %v", err)
			}
			if got.String() != test.want {
				t.Fatalf("setBody() = %q, want %q", got.String(), test.want)
			}
		})
	}
}

func TestSetBodyRequestFile(t *testing.T) {
	file, err := os.CreateTemp(t.TempDir(), "request-body-*")
	if err != nil {
		t.Fatalf("creating request file: %v", err)
	}
	t.Cleanup(func() {
		_ = file.Close()
	})
	if _, err := file.WriteString("file request body"); err != nil {
		t.Fatalf("writing request file: %v", err)
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		t.Fatalf("rewinding request file: %v", err)
	}

	got, err := setBody(file, "application/octet-stream")
	if err != nil {
		t.Fatalf("setBody() error = %v", err)
	}
	if got.String() != "file request body" {
		t.Fatalf("setBody() = %q, want file contents", got.String())
	}
}
