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
	delay = setDelay(req, nil, now)
	assert.Equal(t, time.Duration(0), delay)

	// Test subsequent requests with rate limit data
	// Old concurrent messages should be purged.
	orgRate = RateLimit{
		Limit:         240,
		Remaining:     100,
		Reset:         now.Add(30 * time.Second).Unix(),
		LastRemaining: 101,
		ConcurrentMessages: []time.Time{
			now.Add(-1000 * time.Millisecond),
			now.Add(-750 * time.Millisecond),
			now,
			now.Add(250 * time.Millisecond),
		},
	}
	delay = setDelay(req, nil, now)
	assert.Equal(t, 750*time.Millisecond, delay)

	// All complications from valid state:
	orgRate = RateLimit{
		Limit:         240,
		Remaining:     100,
		Reset:         now.Add(30 * time.Second).Unix(),
		LastRemaining: 104,
		ConcurrentMessages: []time.Time{
			now.Add(1000 * time.Millisecond),
			now.Add(750 * time.Millisecond),
			now.Add(500 * time.Millisecond),
			now.Add(250 * time.Millisecond),
		},
	}
	instantTestRate = orgRate // Use state to test instant test below
	delay = setDelay(req, nil, now)
	assert.Equal(t, 2*time.Second, delay)

	// Same result should be obtained for an instant test
	req, _ = http.NewRequest("GET", "https://api.thousandeyes.com/v6/instant/agent-to-server.json", nil)
	delay = setDelay(req, nil, now)
	assert.Equal(t, 2*time.Second, delay)
	req, _ = http.NewRequest("GET", "https://api.thousandeyes.com/v6/agents.json", nil)

	// LastTime over the minimum delay time should result in the minimum delay time if there
	// are no concurrent messages and last remaining has not decreased by more than 1.
	orgRate.LastRemaining = 101
	orgRate.ConcurrentMessages = []time.Time{}
	delay = setDelay(req, nil, now)
	assert.Equal(t, time.Duration(250*time.Millisecond), delay)

	// A passed response means we should delay, as this is presently only done
	// in response to a 429
	resp.StatusCode = 429
	delay = setDelay(req, resp, now)
	assert.Equal(t, 31*time.Second, delay)

	// Remaining messages being under the minimum should also result in waiting
	// until reset
	orgRate.Remaining = 1
	delay = setDelay(req, nil, now)
	assert.Equal(t, 31*time.Second, delay)

	// Test conflicting or invalid states
	// After reset, LastRemaining may be larger than Remaining
	orgRate = RateLimit{
		Limit:              240,
		Remaining:          240,
		Reset:              now.Add(30 * time.Second).Unix(),
		LastRemaining:      2,
		ConcurrentMessages: []time.Time{},
	}
	delay = setDelay(req, nil, now)
	assert.Equal(t, 250*time.Millisecond, delay)

	// Delays over one minute should be shortened to one minute
	orgRate.Remaining = 0
	orgRate.Reset = now.Add(120 * time.Second).Unix()
	delay = setDelay(req, nil, now)
	assert.Equal(t, 1*time.Minute, delay)

}

func Test_storeLimits(t *testing.T) {
	setup()
	now := time.Now()
	destRate := RateLimit{
		Limit:     240,
		Remaining: 2,
		Reset:     120,
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
