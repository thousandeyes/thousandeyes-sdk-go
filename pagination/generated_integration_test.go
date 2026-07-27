package pagination_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"

	"github.com/thousandeyes/thousandeyes-sdk-go/v3/administrative"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/client"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/endpointtestresults"
)

type recordedRequest struct {
	method      string
	path        string
	query       url.Values
	contentType string
	body        []byte
}

type pageRecorder struct {
	t      *testing.T
	method string
	path   string
	pages  []string
	calls  []recordedRequest
}

func (r *pageRecorder) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		r.t.Errorf("read request body: %v", err)
	}
	r.calls = append(r.calls, recordedRequest{
		method:      request.Method,
		path:        request.URL.Path,
		query:       request.URL.Query(),
		contentType: request.Header.Get("Content-Type"),
		body:        body,
	})
	index := len(r.calls) - 1
	if index >= len(r.pages) {
		r.t.Errorf("unexpected request %d", index+1)
		http.Error(w, "unexpected request", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Page", strconv.Itoa(index+1))
	_, _ = io.WriteString(w, r.pages[index])
}

type contextKey struct{}

type contextObserver struct {
	base http.RoundTripper
	saw  bool
}

func (o *contextObserver) RoundTrip(request *http.Request) (*http.Response, error) {
	if request.Context().Value(contextKey{}) == "wired" {
		o.saw = true
	}
	return o.base.RoundTrip(request)
}

func TestGeneratedGetUserEventsPagination(t *testing.T) {
	recorder := &pageRecorder{
		t:      t,
		method: http.MethodGet,
		path:   "/audit-user-events",
		pages: []string{
			`{"startDate":"2026-01-01T00:00:00Z","_links":{"next":{"href":"https://other.example/wrong?cursor=get-next"}}}`,
			`{"startDate":"2026-01-02T00:00:00Z","_links":{}}`,
			`{"startDate":"2026-01-03T00:00:00Z","_links":{}}`,
		},
	}
	server := httptest.NewServer(recorder)
	defer server.Close()

	observer := &contextObserver{base: server.Client().Transport}
	apiClient := newTestClient(server.URL, observer)
	api := (*administrative.UserEventsAPIService)(&apiClient.Common)
	request := api.GetUserEvents().
		Aid("account").
		UseAllPermittedAids(true).
		Window("1h")
	pager := request.Paginated()
	if len(recorder.calls) != 0 {
		t.Fatal("Paginated performed a request")
	}

	ctx := context.WithValue(context.Background(), contextKey{}, "wired")
	first, firstHTTP, err := pager.NextPage(ctx)
	if err != nil || first == nil || first.StartDate == nil || firstHTTP == nil ||
		firstHTTP.Header.Get("X-Page") != "1" {
		t.Fatalf("first page = (%#v, %#v, %v)", first, firstHTTP, err)
	}
	if !observer.saw {
		t.Fatal("NextPage context was not wired to the HTTP request")
	}
	second, secondHTTP, err := pager.NextPage(context.Background())
	if err != nil || second == nil || second.StartDate == nil || secondHTTP == nil ||
		secondHTTP.Header.Get("X-Page") != "2" {
		t.Fatalf("second page = (%#v, %#v, %v)", second, secondHTTP, err)
	}
	page, response, err := pager.NextPage(context.Background())
	assertCompleted(t, page, response, err)
	assertReplay(t, recorder.calls, http.MethodGet, "/audit-user-events", "get-next")
	if got := recorder.calls[0].query.Encode(); got != "aid=account&useAllPermittedAids=true&window=1h" {
		t.Fatalf("GET query = %q", got)
	}
	legacy, legacyHTTP, err := request.Execute()
	if err != nil || legacy == nil || legacy.StartDate == nil || legacyHTTP == nil ||
		legacyHTTP.Header.Get("X-Page") != "3" {
		t.Fatalf("legacy Execute = (%#v, %#v, %v)", legacy, legacyHTTP, err)
	}
	assertLegacyRequest(t, recorder.calls, http.MethodGet, recorder.path)
}

func TestGeneratedFilterLocalNetworksPagination(t *testing.T) {
	recorder := &pageRecorder{
		t:      t,
		method: http.MethodPost,
		path:   "/endpoint/test-results/local-networks/topologies/filter",
		pages: []string{
			`{"startDate":"2026-01-01T00:00:00Z","_links":{"next":{"href":"https://other.example/wrong?cursor=post-next"}}}`,
			`{"startDate":"2026-01-02T00:00:00Z","_links":{}}`,
			`{"startDate":"2026-01-03T00:00:00Z","_links":{}}`,
		},
	}
	server := httptest.NewServer(recorder)
	defer server.Close()

	apiClient := newTestClient(server.URL, server.Client().Transport)
	api := (*endpointtestresults.LocalNetworkEndpointTestResultsAPIService)(&apiClient.Common)
	filter := endpointtestresults.EndpointNetworkTopologyResultRequestFilter{
		Location:  []string{"Lisbon"},
		NetworkId: []string{"network-1"},
	}
	body := endpointtestresults.EndpointNetworkTopologyResultRequest{SearchFilters: &filter}
	request := api.FilterLocalNetworksTestResultsTopologies().
		Aid("account").
		Window("2h").
		EndpointNetworkTopologyResultRequest(body)
	pager := request.Paginated()
	if len(recorder.calls) != 0 {
		t.Fatal("Paginated performed a request")
	}

	for page := 1; page <= 2; page++ {
		result, response, err := pager.NextPage(context.Background())
		if err != nil || result == nil || result.StartDate == nil ||
			response == nil || response.Header.Get("X-Page") != strconv.Itoa(page) {
			t.Fatalf("page %d = (%#v, %#v, %v)", page, result, response, err)
		}
	}
	page, response, err := pager.NextPage(context.Background())
	assertCompleted(t, page, response, err)
	assertReplay(t, recorder.calls, http.MethodPost, recorder.path, "post-next")
	if got := recorder.calls[0].query.Encode(); got != "aid=account&window=2h" {
		t.Fatalf("POST query = %q", got)
	}

	var firstBody, secondBody, wantBody any
	if err := json.Unmarshal(recorder.calls[0].body, &firstBody); err != nil {
		t.Fatalf("decode first body: %v", err)
	}
	if err := json.Unmarshal(recorder.calls[1].body, &secondBody); err != nil {
		t.Fatalf("decode second body: %v", err)
	}
	if err := json.Unmarshal([]byte(`{"searchFilters":{"location":["Lisbon"],"networkId":["network-1"]}}`), &wantBody); err != nil {
		t.Fatalf("decode expected body: %v", err)
	}
	if !reflect.DeepEqual(firstBody, secondBody) || !reflect.DeepEqual(firstBody, wantBody) ||
		recorder.calls[0].contentType != "application/json" ||
		recorder.calls[1].contentType != "application/json" {
		t.Fatalf("POST body or content type changed between pages")
	}
	legacy, legacyHTTP, err := request.Execute()
	if err != nil || legacy == nil || legacy.StartDate == nil || legacyHTTP == nil ||
		legacyHTTP.Header.Get("X-Page") != "3" {
		t.Fatalf("legacy Execute = (%#v, %#v, %v)", legacy, legacyHTTP, err)
	}
	assertLegacyRequest(t, recorder.calls, http.MethodPost, recorder.path)
}

func newTestClient(serverURL string, transport http.RoundTripper) *client.APIClient {
	configuration := client.NewConfiguration().WithServerUrl(serverURL)
	configuration.HTTPClient = &http.Client{Transport: transport}
	return client.NewAPIClient(configuration)
}

func assertReplay(t *testing.T, calls []recordedRequest, method, path, nextCursor string) {
	t.Helper()
	if len(calls) != 2 {
		t.Fatalf("requests = %d, want 2", len(calls))
	}
	for index := range calls {
		if calls[index].method != method || calls[index].path != path {
			t.Fatalf("request %d = %s %s, want %s %s", index+1, calls[index].method, calls[index].path, method, path)
		}
	}
	if calls[0].query.Get("cursor") != "" || calls[1].query.Get("cursor") != nextCursor {
		t.Fatalf("cursors = %q, %q", calls[0].query.Get("cursor"), calls[1].query.Get("cursor"))
	}
	calls[0].query.Del("cursor")
	calls[1].query.Del("cursor")
	if calls[0].query.Encode() != calls[1].query.Encode() {
		t.Fatalf("non-cursor query changed: %q != %q", calls[0].query.Encode(), calls[1].query.Encode())
	}
}

func assertCompleted[T any](t *testing.T, page *T, response *http.Response, err error) {
	t.Helper()
	if page != nil || response != nil || !errors.Is(err, io.EOF) {
		t.Fatalf("after completion = (%#v, %#v, %v), want (nil, nil, io.EOF)", page, response, err)
	}
}

func assertLegacyRequest(t *testing.T, calls []recordedRequest, method, path string) {
	t.Helper()
	if len(calls) != 3 {
		t.Fatalf("requests after legacy Execute = %d, want 3", len(calls))
	}
	call := calls[2]
	if call.method != method || call.path != path || call.query.Get("cursor") != "" {
		t.Fatalf("legacy request = %s %s?%s", call.method, call.path, call.query.Encode())
	}
}
