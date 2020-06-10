package thousandeyes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
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
