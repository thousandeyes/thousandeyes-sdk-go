package thousandeyes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the PagerDuty client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	var authToken = "foo"
	var accountGroup = "bar"
	client = NewClient(authToken, accountGroup)
}

func teardown() {
	server.Close()
}

func Test_ClientAccountGroup(t *testing.T) {
	setup()
	out := `{"agents": []}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo", AccountGroupID: "test"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "test", r.URL.Query().Get("aid"))
		_, _ = w.Write([]byte(out))
	})
	_, _ = client.GetAgents()
}

func Test_ClientAccountGroupNone(t *testing.T) {
	setup()
	out := `{"agents": []}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/agents.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "", r.URL.Query().Get("aid"))
		_, _ = w.Write([]byte(out))
	})
	_, _ = client.GetAgents()
}

func Test_setDelay(t *testing.T) {
	setup()
	now := time.Now()

	var delay time.Duration
	var req *http.Request
	var resp *http.Response
	orgRate = RateLimit{}
	instantTestRate = RateLimit{}
	req, _ = http.NewRequest("GET", "https://api.thousandeyes.com/v6/agents.json", nil)
	resp = &http.Response{}
	resp.Header = make(map[string][]string)

	// Test initial requests, for which rate limit data is not available
	orgRate = RateLimit{}
	delay = setDelay(req, nil)
	assert.Equal(t, time.Duration(0), delay)

	// Test subsequent requests with rate limit data
	// All complications:
	orgRate = RateLimit{
		Limit:              240,
		Remaining:          100,
		Reset:              now.Add(30 * time.Second).Unix(),
		LastRemaining:      104,
		LastTime:           now.Add(-1 * time.Nanosecond),
		ConcurrentMessages: 3,
	}
	delay = setDelay(req, nil)
	assert.Equal(t, 2*time.Second, delay)

	// LastTime over the minimum delay time should result in the minimum delay time if there
	// are no concurrent messages and last remaining has not decreased by more than 1.
	orgRate.LastRemaining = 101
	orgRate.ConcurrentMessages = 0
	orgRate.LastTime = now.Add(-5 * time.Second)
	delay = setDelay(req, nil)
	assert.Equal(t, time.Duration(250*time.Millisecond), delay)

	// A passed response means we should delay, as this is presently only done
	// in response to a 429
	resp.StatusCode = 429
	delay = setDelay(req, resp)
	assert.Equal(t, 31*time.Second, delay)

	// Remaining messages being under the minimum should also result in waiting
	// until reset
	orgRate.Remaining = 1
	delay = setDelay(req, nil)
	assert.Equal(t, 31*time.Second, delay)

}

func Test_storeLimits(t *testing.T) {
	setup()
	now := time.Now()
	destRate := RateLimit{
		Limit:     240,
		Remaining: 2,
		Reset:     120,
		LastTime:  now,
	}

	var req *http.Request
	var resp *http.Response
	orgRate = RateLimit{}
	instantTestRate = RateLimit{}
	req, _ = http.NewRequest("GET", "https://api.thousandeyes.com/v6/agents.json", nil)
	resp = &http.Response{}
	resp.Header = make(map[string][]string)
	resp.Header.Add("X-Organization-Rate-Limit-Limit", "240")
	resp.Header.Add("X-Organization-Rate-Limit-Remaining", "2")
	resp.Header.Add("X-Organization-Rate-Limit-Reset", "120")
	storeLimits(req, resp, now)
	assert.Equal(t, destRate, orgRate)
	assert.Equal(t, RateLimit{}, instantTestRate)

	orgRate = RateLimit{}
	instantTestRate = RateLimit{}
	req, _ = http.NewRequest("GET", "https://api.thousandeyes.com/v6/instant/agent-to-server.json", nil)
	resp = &http.Response{}
	resp.Header = make(map[string][]string)
	resp.Header.Add("X-Instant-Test-Rate-Limit-Limit", "240")
	resp.Header.Add("X-Instant-Test-Rate-Limit-Remaining", "2")
	resp.Header.Add("X-Instant-Test-Rate-Limit-Reset", "120")
	storeLimits(req, resp, now)
	assert.Equal(t, destRate, instantTestRate)
	assert.Equal(t, RateLimit{}, orgRate)
}

func Test_isInstantTest(t *testing.T) {
	var req *http.Request
	req, _ = http.NewRequest("GET", "https://api.thousandeyes.com/v6/instant/agent-to-server.json", nil)
	assert.Equal(t, true, isInstantTest(req))
	req, _ = http.NewRequest("GET", "https://api.thousandeyes.com/v6/endpoint-instant/agent-to-server.json", nil)
	assert.Equal(t, true, isInstantTest(req))
	req, _ = http.NewRequest("GET", "https://api.thousandeyes.com/v6/agents.json", nil)
	assert.Equal(t, false, isInstantTest(req))
}
